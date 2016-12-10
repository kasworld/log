// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// python like log package
// can use instead of standard log package
package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Log struct {
	mutex    sync.Mutex // ensures atomic writes; protects the following fields
	flag     LF_Type    // properties
	loglevel LL_Type
	prefix   string // prefix to write at beginning of each line
	filename string
	out      io.WriteCloser
}

func New(w io.WriteCloser, prefix string, loglevel LL_Type, release bool) *Log {
	flags := LF_stdFlags
	if !release {
		flags = LF_time | LF_shortfile | LF_functionname
	}

	l := Log{
		prefix:   prefix,
		flag:     flags,
		loglevel: loglevel,
		out:      w,
	}
	return &l
}

func NewFile(filename string, prefix string, loglevel LL_Type, release bool) (*Log, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	l := New(f, prefix, loglevel, release)
	l.filename = filename
	return l, nil
}

func (l *Log) Reload() error {
	if l.filename == "" {
		return fmt.Errorf("not file log")
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.out.Close()

	out, err := os.OpenFile(l.filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	l.out = out
	return nil
}

func (l Log) LogPrintf(calldepth int, ll LL_Type, format string, v ...interface{}) []byte {
	if !l.IsLevel(ll) {
		return nil
	}
	llinfo := fmt.Sprintf("%s", ll)
	s := fmt.Sprintf(format, v...)

	var buf []byte
	FormatHeader(&buf, calldepth+1, time.Now(), l.flag, l.prefix, llinfo)
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return buf
}

func (l *Log) Output(s []byte) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	_, err := l.out.Write(s)
	return err
}
