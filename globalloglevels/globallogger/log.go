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

// global log instance
package globallogger

import (
	"github.com/kasworld/log/globalloglevels"
	"github.com/kasworld/log/globalloglevels/globallogbase"
	"github.com/kasworld/log/logflagi"
	"github.com/kasworld/log/logflags"
)

var GlobalLogger = globallogbase.New("",
	logflags.DefaultValue(false), globalloglevels.LL_All)

func init() {
	GlobalLogger.AddDestination(globalloglevels.LL_All, globallogbase.OutputStderr)
}

func Reload() error {
	return GlobalLogger.Reload()
}

func GetLogger() *globallogbase.LogBase {
	return GlobalLogger
}
func SetLogger(l *globallogbase.LogBase) {
	GlobalLogger = l
}

func LevelString() string {
	return GlobalLogger.String()
}

func AddLevel(level globalloglevels.LL_Type) {
	GlobalLogger.AddLevel(level)
}

func SetLevel(level globalloglevels.LL_Type) {
	GlobalLogger.SetLevel(level)
}

func DelLevel(level globalloglevels.LL_Type) {
	GlobalLogger.DelLevel(level)
}

func IsLevel(level globalloglevels.LL_Type) bool {
	return GlobalLogger.IsLevel(level)
}

func SetPrefix(p string) {
	GlobalLogger.SetPrefix(p)
}

// Prefix returns the output prefix for the GlobalLogger.
func GetPrefix() string {
	return GlobalLogger.GetPrefix()
}

// Flags returns the output flags for the GlobalLogger.
func GetFlags() logflagi.LogFlagI {
	return GlobalLogger.GetFlags()
}

// SetFlags sets the output flags for the GlobalLogger.
func SetFlags(flag logflagi.LogFlagI) {
	GlobalLogger.SetFlags(flag)
}

func Panic(format string, v ...interface{}) error {
	s, err := GlobalLogger.LogPrintf(2, globalloglevels.LL_Fatal, format, v...)
	panic(string(s))
	return err
}
