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
	"sync"
	"time"

	"github.com/kasworld/log/logdest_file"
	"github.com/kasworld/log/logdesti"
	"github.com/kasworld/log/logflags"
	"github.com/kasworld/log/logformater"
	"github.com/kasworld/log/loglevels"
)

type Log struct {
	mutex    sync.Mutex       // ensures atomic writes; protects the following fields
	flag     logflags.LF_Type // properties
	loglevel loglevels.LL_Type
	prefix   string // prefix to write at beginning of each line
	logdst   logdesti.LogDestI
	// filename string
	// out      io.WriteCloser
}

func New(dst logdesti.LogDestI, prefix string, loglevel loglevels.LL_Type, release bool) *Log {
	flags := logflags.LF_stdFlags
	if !release {
		flags = logflags.LF_time | logflags.LF_shortfile | logflags.LF_functionname
	}

	l := Log{
		prefix:   prefix,
		flag:     flags,
		loglevel: loglevel,
		logdst:   dst,
	}
	return &l
}

func NewFile(filename string, prefix string, loglevel loglevels.LL_Type, release bool) (*Log, error) {
	dst, err := logdest_file.New(filename)
	if err != nil {
		return nil, err
	}
	l := New(dst, prefix, loglevel, release)
	return l, nil
}

func (l *Log) Reload() error {
	return l.logdst.Reload()
}

func (l Log) LogPrintf(calldepth int, ll loglevels.LL_Type, format string, v ...interface{}) []byte {
	if !l.IsLevel(ll) {
		return nil
	}
	s := fmt.Sprintf(format, v...)

	var buf []byte
	logformater.FormatHeader(&buf, calldepth+1, time.Now(), l.flag, l.prefix, ll)
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return buf
}

func (l *Log) Output(b []byte) error {
	return l.logdst.Output(b)
}
