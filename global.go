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
	logger.printf(LL_Info, format, v...)
}
func Info(format string, v ...interface{}) {
	logger.printf(LL_Info, format, v...)
}
func Warn(format string, v ...interface{}) {
	logger.printf(LL_Warn, format, v...)
}
func Debug(format string, v ...interface{}) {
	logger.printf(LL_Debug, format, v...)
}
func Error(format string, v ...interface{}) {
	logger.printf(LL_Error, format, v...)
}
func Fatal(format string, v ...interface{}) {
	logger.printf(LL_Fatal, format, v...)
	os.Exit(1)
}

func AddLevel(level int) {
	logger.loglevel |= level
}
func SetLevel(level int) {
	logger.loglevel = level
}
func DelLevel(level int) {
	logger.loglevel &= ^level
}
func IsLevel(level int) bool {
	return logger.loglevel&level != 0
}

func SetPrefix(p string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.prefix = p
}

// Prefix returns the output prefix for the logger.
func GetPrefix() string {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.prefix
}

// Flags returns the output flags for the logger.
func GetFlags() int {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.flag
}

// SetFlags sets the output flags for the logger.
func SetFlags(flag int) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.flag = flag
}
