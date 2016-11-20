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
	"runtime"
	"sync"
	"time"
)

type Log struct {
	mutex    sync.Mutex // ensures atomic writes; protects the following fields
	flag     LF_Type    // properties
	prefix   string     // prefix to write at beginning of each line
	loglevel LL_Type
	out      io.WriteCloser
	filename string
	buf      []byte // for accumulating text to write
}

func New(w io.WriteCloser, prefix string, loglevel LL_Type, release bool) *Log {
	flags := LstdFlags
	if !release {
		flags = Ltime | Lshortfile | Lfunctionname
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

func (l *Log) Output(s []byte) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	_, err := l.out.Write(s)
	return err
}

func (l Log) LogPrintf(calldepth int, ll LL_Type, format string, v ...interface{}) []byte {
	if !l.IsLevel(ll) {
		return nil
	}
	llinfo := fmt.Sprintf("%s", ll)
	s := fmt.Sprintf(format, v...)

	// log.output
	now := time.Now() // get this early.
	var file string
	var fnname string
	var line int
	var pc uintptr
	if l.flag&(Lshortfile|Llongfile|Lfunctionname) != 0 {
		// release lock while getting caller info - it's expensive.
		var ok bool
		pc, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
			fnname = "???"
		} else if l.flag&Lfunctionname != 0 {
			fn := runtime.FuncForPC(pc)
			fnname = fn.Name()
		}
	}
	var buf []byte
	l.formatHeader(&buf, llinfo, now, file, line, fnname)
	buf = append(buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return buf
}

func (l Log) formatHeader(
	buf *[]byte,
	llinfo string,
	t time.Time,
	file string,
	line int,
	fnname string) {

	if l.flag&Lprefix != 0 {
		*buf = append(*buf, l.prefix...)
		*buf = append(*buf, ' ')
	}
	*buf = append(*buf, llinfo...)
	*buf = append(*buf, ' ')
	if l.flag&LUTC != 0 {
		t = t.UTC()
	}
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(Lshortfile|Llongfile) != 0 {
		if l.flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		if l.flag&(Lfunctionname) != 0 {
			*buf = append(*buf, ':')
			*buf = append(*buf, fnname...)
		}
		*buf = append(*buf, ": "...)
	}
}
