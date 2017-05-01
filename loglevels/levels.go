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

package loglevels

import (
	"bytes"
	"fmt"
)

//go:generate stringer -type=LL_Type
type LL_Type int

const (
	LL_Fatal LL_Type = 1 << iota
	LL_Error
	LL_Warn
	LL_Debug
	LL_Info
	LL_END
	LL_All = LL_Debug | LL_Info | LL_Warn | LL_Error | LL_Fatal
)

func (ll LL_Type) LevelString() string {
	var buff bytes.Buffer
	buff.WriteString("LogLevel[")
	for i := LL_Type(1); i < LL_END; i <<= 1 {
		if ll&i != 0 {
			fmt.Fprintf(&buff, "%s ", i)
		}
	}
	buff.WriteString("]")
	return buff.String()
}

func AllLevelString() string {
	var buff bytes.Buffer
	buff.WriteString("LogLevel[")
	for i := LL_Type(1); i < LL_END; i <<= 1 {
		fmt.Fprintf(&buff, "%s ", i)
	}
	buff.WriteString("]")
	return buff.String()
}

func (l *LL_Type) AddLevel(level LL_Type) {
	*l |= level
}
func (l *LL_Type) SetLevel(level LL_Type) {
	*l = level
}
func (l *LL_Type) DelLevel(level LL_Type) {
	*l &= ^level
}
func (l LL_Type) IsLevel(level LL_Type) bool {
	return l&level != 0
}
