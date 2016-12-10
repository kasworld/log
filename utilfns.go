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
	"runtime"
	"time"
)

//// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func FormatHeader(
	buf *[]byte,
	calldepth int,
	now time.Time,
	logflag LF_Type, prefix string, llinfo string) {

	// now := time.Now() // get this early.
	var file string
	var fnname string
	var line int
	var pc uintptr
	if logflag&(LF_shortfile|LF_longfile|LF_functionname) != 0 {
		// release lock while getting caller info - it's expensive.
		var ok bool
		pc, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
			fnname = "???"
		} else if logflag&LF_functionname != 0 {
			fn := runtime.FuncForPC(pc)
			fnname = fn.Name()
		}
	}

	if logflag&LF_prefix != 0 {
		*buf = append(*buf, prefix...)
		*buf = append(*buf, ' ')
	}
	*buf = append(*buf, llinfo...)
	*buf = append(*buf, ' ')
	if logflag&LF_UTC != 0 {
		now = now.UTC()
	}
	if logflag&(LF_date|LF_time|LF_microseconds) != 0 {
		if logflag&LF_date != 0 {
			year, month, day := now.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if logflag&(LF_time|LF_microseconds) != 0 {
			hour, min, sec := now.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if logflag&LF_microseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, now.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if logflag&(LF_shortfile|LF_longfile) != 0 {
		if logflag&LF_shortfile != 0 {
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
		if logflag&(LF_functionname) != 0 {
			*buf = append(*buf, ':')
			*buf = append(*buf, fnname...)
		}
		*buf = append(*buf, ": "...)
	}
}
