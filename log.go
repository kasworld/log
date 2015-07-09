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
	"log"
	"os"
)

const (
	LL_Debug = 1 << iota
	LL_Info
	LL_Warn
	LL_Error
	LL_Fatal
	LL_All = LL_Debug | LL_Info | LL_Warn | LL_Error | LL_Fatal
)

var LogLevel2str = map[int]string{
	LL_Debug: "Debug",
	LL_Info:  "Info",
	LL_Warn:  "Warn",
	LL_Error: "Error",
	LL_Fatal: "Fatal",
}

type Log struct {
	loglevel int
	l        map[int]*log.Logger
}

func New(prefix string, loglevel int, release bool) *Log {
	l := Log{
		loglevel: loglevel,
		l:        make(map[int]*log.Logger),
	}
	flags := log.LstdFlags
	if !release {
		flags = log.Ltime | log.Lshortfile
	}
	for i, v := range LogLevel2str {
		l.l[i] = log.New(os.Stderr, fmt.Sprintf("%v %v:", prefix, v), flags)
	}
	return &l
}

func (l *Log) SetPrefix(p string) {
	for i, _ := range LogLevel2str {
		l.l[i].SetPrefix(p)
	}
}

func (l Log) String() string {
	levelstr := ""
	for i, v := range LogLevel2str {
		if l.IsLevel(i) {
			levelstr += v + " "
		}
	}
	return fmt.Sprintf("log level is %v", levelstr)
}
func (l *Log) AddLevel(level int) {
	l.loglevel |= level
}
func (l *Log) SetLevel(level int) {
	l.loglevel = level
}
func (l *Log) DelLevel(level int) {
	l.loglevel &= ^level
}
func (l *Log) IsLevel(level int) bool {
	return l.loglevel&level != 0
}

func (l Log) printf(ll int, format string, v ...interface{}) {
	if !l.IsLevel(ll) {
		return
	}
	l.l[ll].Output(3, fmt.Sprintf(format, v...))
}

func (l Log) Info(format string, v ...interface{}) {
	l.printf(LL_Info, format, v...)
}
func (l Log) Warn(format string, v ...interface{}) {
	l.printf(LL_Warn, format, v...)
}
func (l Log) Debug(format string, v ...interface{}) {
	l.printf(LL_Debug, format, v...)
}
func (l Log) Error(format string, v ...interface{}) {
	l.printf(LL_Error, format, v...)
}
func (l Log) Fatal(format string, v ...interface{}) {
	l.printf(LL_Fatal, format, v...)
	os.Exit(1)
}

// ===

var logger = New("Global", LL_All, false)

func LevelString() string {
	return logger.String()
}

func SetReleaseLogger() {
	logger = New("", LL_All, true)
}
func Printf(format string, v ...interface{}) {
	logger.printf(LL_Info, format, v...)
}
func Fatalf(format string, v ...interface{}) {
	logger.printf(LL_Fatal, format, v...)
	os.Exit(1)
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
