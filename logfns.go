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
	"bytes"
	"fmt"
	"os"
)

func (l Log) String() string {
	var buff bytes.Buffer
	buff.WriteString("Log[")
	for i := LL_Type(1); i < LL_END; i <<= 1 {
		fmt.Fprintf(&buff, "%s ", i)
	}
	buff.WriteString("]")
	return buff.String()
}

func (l *Log) SetPrefix(p string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.prefix = p
}

// Prefix returns the output prefix for the logger.
func (l Log) GetPrefix() string {
	return l.prefix
}

// Flags returns the output flags for the logger.
func (l Log) GetFlags() LF_Type {
	return l.flag
}

// SetFlags sets the output flags for the logger.
func (l *Log) SetFlags(flag LF_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.flag = flag
}

func (l *Log) AddLevel(level LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel |= level
}
func (l *Log) SetLevel(level LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel = level
}
func (l *Log) DelLevel(level LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel &= ^level
}
func (l Log) IsLevel(level LL_Type) bool {
	return l.loglevel&level != 0
}

func (l *Log) Printf(ll LL_Type, format string, v ...interface{}) error {
	// calldepth := 2
	s := l.LogPrintf(2, ll, format, v...)
	return l.Output(s)
}

func (l *Log) Info(format string, v ...interface{}) {
	s := l.LogPrintf(2, LL_Info, format, v...)
	l.Output(s)
}
func (l *Log) Warn(format string, v ...interface{}) {
	s := l.LogPrintf(2, LL_Warn, format, v...)
	l.Output(s)
}
func (l *Log) Debug(format string, v ...interface{}) {
	s := l.LogPrintf(2, LL_Debug, format, v...)
	l.Output(s)
}
func (l *Log) Error(format string, v ...interface{}) {
	s := l.LogPrintf(2, LL_Error, format, v...)
	l.Output(s)
}
func (l *Log) Fatal(format string, v ...interface{}) {
	s := l.LogPrintf(2, LL_Fatal, format, v...)
	l.Output(s)
	os.Exit(1)
}
