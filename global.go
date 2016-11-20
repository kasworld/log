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
package log

import (
	// "io"
	"os"
)

var logger = New(os.Stdout, "", LL_All, false)

func SetLogger(l *Log) {
	logger = l
}

func Reload() error {
	return logger.Reload()
}

func LevelString() string {
	return logger.String()
}

func Printf(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Info, format, v...)
	logger.Output(s)
}
func Info(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Info, format, v...)
	logger.Output(s)
}
func Warn(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Warn, format, v...)
	logger.Output(s)
}
func Debug(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Debug, format, v...)
	logger.Output(s)
}
func Error(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Error, format, v...)
	logger.Output(s)
}
func Fatal(format string, v ...interface{}) {
	s := logger.LogPrintf(2, LL_Fatal, format, v...)
	logger.Output(s)
	os.Exit(1)
}

func AddLevel(level LL_Type) {
	logger.AddLevel(level)
}
func SetLevel(level LL_Type) {
	logger.SetLevel(level)
}
func DelLevel(level LL_Type) {
	logger.DelLevel(level)
}
func IsLevel(level LL_Type) bool {
	return logger.IsLevel(level)
}

func SetPrefix(p string) {
	logger.SetPrefix(p)
}

// Prefix returns the output prefix for the logger.
func GetPrefix() string {
	return logger.GetPrefix()
}

// Flags returns the output flags for the logger.
func GetFlags() LF_Type {
	return logger.GetFlags()
}

// SetFlags sets the output flags for the logger.
func SetFlags(flag LF_Type) {
	logger.SetFlags(flag)
}
