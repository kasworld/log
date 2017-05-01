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

import "github.com/kasworld/log/loglevels"

func (l *LogBase) Info(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Info, format, v...)
	l.Output(s)
}
func (l *LogBase) Warn(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Warn, format, v...)
	l.Output(s)
}
func (l *LogBase) Debug(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Debug, format, v...)
	l.Output(s)
}
func (l *LogBase) Error(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Error, format, v...)
	l.Output(s)
}

func (l *LogBase) Fatal(format string, v ...interface{}) {
	s := l.LogPrintf(2, loglevels.LL_Fatal, format, v...)
	l.Output(s)
}
