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

package globallogbase

import (
	"fmt"

	"github.com/kasworld/log/globalloglevels"
	"github.com/kasworld/log/logflagi"
)

func (lg *LogBase) AddLevel(level globalloglevels.LL_Type) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.loglevel |= level
}

func (lg *LogBase) SetLevel(level globalloglevels.LL_Type) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.loglevel = level
}

func (lg *LogBase) DelLevel(level globalloglevels.LL_Type) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.loglevel &= ^level
}

func (lg *LogBase) IsLevel(level globalloglevels.LL_Type) bool {
	return lg.loglevel&level != 0
}

func (lg *LogBase) FlagString() string {
	return lg.flag.FlagString()
}

func (lg *LogBase) LevelString() string {
	return lg.loglevel.LevelsString()
}

func (lg *LogBase) SetPrefix(p string) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.prefix = p
}

// Prefix returns the output prefix for the logger.
func (lg *LogBase) GetPrefix() string {
	return lg.prefix
}

// Flags returns the output flags for the logger.
func (lg *LogBase) GetFlags() logflagi.LogFlagI {
	return lg.flag
}

// SetFlags sets the output flags for the logger.
func (lg *LogBase) SetFlags(flag logflagi.LogFlagI) {
	lg.mutex.Lock()
	defer lg.mutex.Unlock()
	lg.flag = flag
}

func (lg LogBase) String() string {
	return fmt.Sprintf("LogBase[%v %v]",
		lg.FlagString(), lg.LevelString(),
	)
}
