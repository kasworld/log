// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package basiclog

import (
	"fmt"
)

const (
	LL_Fatal LL_Type = 1 << iota
	LL_Error
	LL_Warn
	LL_Debug
	LL_END

	LL_All = LL_END - 1
)

var leveldata = map[LL_Type]string{
	1:  "Fatal",
	2:  "Error",
	4:  "Warn",
	8:  "Debug",
	16: "END",
}

//////////////////////////////////////////////////////////////////////

func (l *LogBase) Fatal(format string, v ...interface{}) {
	s := l.Format2Bytes(2, LL_Fatal, format, v...)
	err := l.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithFatalLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, LL_Fatal, format, v...)
	err := l.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	s := GlobalLogger.Format2Bytes(2, LL_Fatal, format, v...)
	err := GlobalLogger.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
}
func NewErrorWithFatalLog(format string, v ...interface{}) error {
	s := GlobalLogger.Format2Bytes(2, LL_Fatal, format, v...)
	err := GlobalLogger.Output(LL_Fatal, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Error(format string, v ...interface{}) {
	s := l.Format2Bytes(2, LL_Error, format, v...)
	err := l.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithErrorLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, LL_Error, format, v...)
	err := l.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}
func Error(format string, v ...interface{}) {
	s := GlobalLogger.Format2Bytes(2, LL_Error, format, v...)
	err := GlobalLogger.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
}
func NewErrorWithErrorLog(format string, v ...interface{}) error {
	s := GlobalLogger.Format2Bytes(2, LL_Error, format, v...)
	err := GlobalLogger.Output(LL_Error, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Warn(format string, v ...interface{}) {
	s := l.Format2Bytes(2, LL_Warn, format, v...)
	err := l.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithWarnLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, LL_Warn, format, v...)
	err := l.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}
func Warn(format string, v ...interface{}) {
	s := GlobalLogger.Format2Bytes(2, LL_Warn, format, v...)
	err := GlobalLogger.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
}
func NewErrorWithWarnLog(format string, v ...interface{}) error {
	s := GlobalLogger.Format2Bytes(2, LL_Warn, format, v...)
	err := GlobalLogger.Output(LL_Warn, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}

func (l *LogBase) Debug(format string, v ...interface{}) {
	s := l.Format2Bytes(2, LL_Debug, format, v...)
	err := l.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
}
func (l *LogBase) NewErrorWithDebugLog(format string, v ...interface{}) error {
	s := l.Format2Bytes(2, LL_Debug, format, v...)
	err := l.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}
func Debug(format string, v ...interface{}) {
	s := GlobalLogger.Format2Bytes(2, LL_Debug, format, v...)
	err := GlobalLogger.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
}
func NewErrorWithDebugLog(format string, v ...interface{}) error {
	s := GlobalLogger.Format2Bytes(2, LL_Debug, format, v...)
	err := GlobalLogger.Output(LL_Debug, s)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Errorf(format, v...)
}
