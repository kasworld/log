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
package logbase

import (
	"fmt"

	"github.com/kasworld/log/logflags"
	"github.com/kasworld/log/loglevels"
)

func (l Log) String() string {
	return fmt.Sprintf("Log[%v, %v]",
		l.flag.FlagString(),
		l.loglevel.LevelString())
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
func (l Log) GetFlags() logflags.LF_Type {
	return l.flag
}

// SetFlags sets the output flags for the logger.
func (l *Log) SetFlags(flag logflags.LF_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.flag = flag
}

func (l *Log) AddLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel |= level
}
func (l *Log) SetLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel = level
}
func (l *Log) DelLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel &= ^level
}
func (l Log) IsLevel(level loglevels.LL_Type) bool {
	return l.loglevel&level != 0
}

func (l *Log) Printf(ll loglevels.LL_Type, format string, v ...interface{}) error {
	s := l.LogPrintf(2, ll, format, v...)
	return l.Output(s)
}