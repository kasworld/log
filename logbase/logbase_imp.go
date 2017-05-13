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
	"os"

	"github.com/kasworld/log/logflags"
	"github.com/kasworld/log/loglevels"
)

func (l LogBase) String() string {
	return fmt.Sprintf("LogBase[%v, %v]",
		l.flag.FlagString(),
		l.loglevel.LevelString())
}

func (l *LogBase) SetPrefix(p string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.prefix = p
}

// Prefix returns the output prefix for the logger.
func (l LogBase) GetPrefix() string {
	return l.prefix
}

// Flags returns the output flags for the logger.
func (l LogBase) GetFlags() logflags.LF_Type {
	return l.flag
}

// SetFlags sets the output flags for the logger.
func (l *LogBase) SetFlags(flag logflags.LF_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.flag = flag
}

func (l *LogBase) AddLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel.AddLevel(level)
}
func (l *LogBase) SetLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel.SetLevel(level)
}
func (l *LogBase) DelLevel(level loglevels.LL_Type) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.loglevel.DelLevel(level)
}
func (l LogBase) IsLevel(level loglevels.LL_Type) bool {
	return l.loglevel.IsLevel(level)
}

func (l *LogBase) Printf(ll loglevels.LL_Type, format string, v ...interface{}) error {
	s := l.LogPrintf(2, ll, format, v...)
	return l.Output(s)
}

func (l *LogBase) Panic(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Fatal, format, v...)
	l.Output(s)
	os.Exit(1)
}

func (l *LogBase) NewErrorWithLog(
	ll loglevels.LL_Type, format string, v ...interface{}) error {

	s := l.LogPrintf(2, ll, format, v...)
	l.Output(s)
	return fmt.Errorf(format, v...)
}
