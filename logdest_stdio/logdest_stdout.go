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

package logdest_stdio

import (
	"os"

	"github.com/kasworld/log/logdesti"
)

type LogDestStdOut struct {
}

func NewStdOut() logdesti.LogDestI {
	ldst := &LogDestStdOut{}
	return ldst
}

func (ldst *LogDestStdOut) GetName() string {
	return os.Stdout.Name()
}
func (ldst *LogDestStdOut) Reload() error {
	return os.Stdout.Sync()

}
func (ldst *LogDestStdOut) Output(b []byte) error {
	_, err := os.Stdout.Write(b)
	return err
}
func (ldst *LogDestStdOut) Flush() error {
	return os.Stdout.Sync()
}
