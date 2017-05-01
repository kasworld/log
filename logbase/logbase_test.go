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

import (
	"testing"

	"github.com/kasworld/log/logdest_stdio"
	"github.com/kasworld/log/loglevels"
)

func TestLogBase_AddLevel(t *testing.T) {
	logger := New(logdest_stdio.NewStdOut(), "", loglevels.LL_All, false)
	t.Logf("%v", logger)

	logger.DelLevel(loglevels.LL_Debug)
	t.Logf("%v", logger)
}
