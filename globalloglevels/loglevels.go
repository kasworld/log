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

// log level
package globalloglevels

import (
	"bytes"
	"fmt"
	"math/bits"
)

//go:generate stringer -type=LL_Type
type LL_Type uint64

func (ll LL_Type) ToShiftedNum() int {
	return bits.Len(uint(ll)) - 1
}

const (
	LL_Fatal LL_Type = 1 << iota
	LL_Error
	LL_Warn
	LL_TraceService // start , load , init , stop, flush , end
	LL_Monitor      // routine info log
	LL_Debug        // only for debug
	LL_AdminAudit   // for web admin action audit
	LL_Analysis
	LL_TraceUser   // major user action tracking , login , logout, kick, etc
	LL_TraceClient // only for generated code and client
	LL_TraceAO
	LL_TraceAI
	LL_TraceTask
	LL_TraceSuspect
	LL_TraceRPC // only for generated code, server to server
	LL_END
	// preset
	LL_All            = LL_END - 1
	LL_ServiceDefault = LL_Fatal | LL_Error | LL_Warn | LL_TraceService | LL_TraceUser | LL_AdminAudit
	LL_TestDefault    = LL_Fatal | LL_Error | LL_Warn | LL_TraceService | LL_TraceUser | LL_Debug | LL_AdminAudit
)

func (ll LL_Type) LevelsString() string {
	var buff bytes.Buffer

	buff.WriteString("LL_Type[")
	for i := LL_Type(1); i < LL_END; i <<= 1 {
		if ll.IsLevel(i) {
			fmt.Fprintf(&buff, "%s, ", i)
		}
	}
	buff.WriteString("]")
	return buff.String()
}

func (ll LL_Type) IsLevel(level LL_Type) bool {
	return ll&level != 0
}
