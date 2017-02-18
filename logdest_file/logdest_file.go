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

package logdest_file

import (
	"fmt"
	"os"
	"sync"

	"github.com/kasworld/log/logdesti"
)

type LogDestFile struct {
	mutex    sync.Mutex // ensures atomic writes; protects the following fields
	filename string
	fd       *os.File
}

func New(filename string) (logdesti.LogDestI, error) {
	ldf := &LogDestFile{}
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	ldf.fd = fd
	ldf.filename = filename
	return ldf, nil
}

func (ldf *LogDestFile) GetName() string {
	return ldf.filename
}
func (ldf *LogDestFile) Reload() error {
	ldf.mutex.Lock()
	defer ldf.mutex.Unlock()
	if ldf.filename == "" {
		return fmt.Errorf("not file log")
	}
	ldf.fd.Close()

	fd, err := os.OpenFile(ldf.filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	ldf.fd = fd
	return nil
}
func (ldf *LogDestFile) Output(b []byte) error {
	ldf.mutex.Lock()
	defer ldf.mutex.Unlock()
	_, err := ldf.fd.Write(b)
	return err

}
func (ldf *LogDestFile) Flush() error {
	return ldf.fd.Sync()
}
