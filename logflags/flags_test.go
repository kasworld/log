// Copyright 2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
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

package logflags

import (
	"testing"
	"time"
)

func TestLF_Type_FormatHeader(t *testing.T) {
	lf := DefaultValue(false)

	var buf []byte
	lf.FormatHeader(&buf, 0, time.Now(), "prefix", "llinfo")
	t.Logf("%v", string(buf))
}

func TestLF_Type_ParseHeader(t *testing.T) {
	lf := DefaultValue(false)

	var buf []byte
	lf.FormatHeader(&buf, 0, time.Now(), "prefix", "llinfo")
	t.Logf("%v", string(buf))
	prefix, llinfo, datestr, timestr, filestr, remain := lf.ParseHeader(buf)
	t.Logf("%v %v %v %v %v %v",
		prefix, llinfo, datestr, timestr, filestr, remain)
}
