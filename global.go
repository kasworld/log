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

	"fmt"

	"github.com/kasworld/log/logbase"
	"github.com/kasworld/log/logdest_stdio"
	"github.com/kasworld/log/logflags"
	"github.com/kasworld/log/loglevels"
)

var logger = logbase.New(logdest_stdio.NewStdOut(), "", loglevels.LL_All, false)

func SetLogger(l *logbase.LogBase) {
	logger = l
}

func Reload() error {
	return logger.Reload()
}

func LevelString() string {
	return logger.String()
}

func Printf(level loglevels.LL_Type, format string, v ...interface{}) {
	s := logger.LogPrintf(2, level, format, v...)
	logger.Output(s)
}

func NewErrorWithLog(level loglevels.LL_Type, format string, v ...interface{}) error {
	s := logger.LogPrintf(2, level, format, v...)
	logger.Output(s)
	return fmt.Errorf(format, v...)
}

func AddLevel(level loglevels.LL_Type) {
	logger.AddLevel(level)
}
func SetLevel(level loglevels.LL_Type) {
	logger.SetLevel(level)
}
func DelLevel(level loglevels.LL_Type) {
	logger.DelLevel(level)
}
func IsLevel(level loglevels.LL_Type) bool {
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
func GetFlags() logflags.LF_Type {
	return logger.GetFlags()
}

// SetFlags sets the output flags for the logger.
func SetFlags(flag logflags.LF_Type) {
	logger.SetFlags(flag)
}
