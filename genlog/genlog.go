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

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	packagename    = flag.String("packagename", "", "packagename")
	leveldatafile  = flag.String("leveldatafile", "", "leveldatafile")
	outputfilename = "log_gen.go"
)

func MakeGenComment() string {
	return fmt.Sprintf(
		"// Code generated by \"%s %s\" \n",
		filepath.Base(os.Args[0]),
		strings.Join(os.Args[1:], " "))

}

func isDir(path string) error {
	finfo, staterr := os.Stat(path)
	if staterr != nil {
		return staterr
	}
	if !finfo.IsDir() {
		return fmt.Errorf("invalid dir: %v", path)
	}
	return nil
}

// loadEnumWithComment load list of enum + comment
func loadEnumWithComment(filename string) ([][]string, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	rtn := make([][]string, 0)
	rd := bufio.NewReader(fd)
	for {
		line, err := rd.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			s2 := strings.SplitN(line, " ", 2)
			if len(s2) == 1 {
				s2 = append(s2, "")
			}
			rtn = append(rtn, s2)
		}
		if err != nil { // eof
			break
		}
	}
	return rtn, nil
}

func SaveTo(outdata *bytes.Buffer, outfilename string) error {
	src, err := format.Source(outdata.Bytes())
	if err != nil {
		fmt.Println(outdata)
		return err
	}
	if werr := ioutil.WriteFile(outfilename, src, 0644); werr != nil {
		return werr
	}
	return nil
}

func main() {
	flag.Parse()

	if *packagename == "" {
		fmt.Printf("packagename not set\n")
		return
	}
	if *leveldatafile == "" {
		fmt.Printf("leveldatafile not set\n")
		return
	}

	leveldata, err := loadEnumWithComment(*leveldatafile)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	buf, err := Build(*packagename, leveldata)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	if err := SaveTo(buf, *packagename+"/"+outputfilename); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func Build(packagename string, leveldata [][]string) (*bytes.Buffer, error) {

	var buff bytes.Buffer

	fmt.Fprintln(&buff, MakeGenComment())
	fmt.Fprintf(&buff, `
		package %[1]s
		import "fmt"
		`,
		packagename)

	fmt.Fprintf(&buff, "const (\n")
	for i, lvname := range leveldata {
		if i == 0 {
			fmt.Fprintf(&buff, "LL_%v LL_Type = 1 << iota // %v\n", lvname[0], lvname[1])
		} else {
			fmt.Fprintf(&buff, "LL_%v // %v\n", lvname[0], lvname[1])
		}
	}
	fmt.Fprintf(&buff, `
	LL_END
	LL_All = LL_END - 1
	LL_Count = %v
	)`, len(leveldata))

	fmt.Fprintf(&buff, `
	var leveldata = map[LL_Type]string{
	`)

	for i, lvname := range leveldata {
		fmt.Fprintf(&buff, "%v : \"%v\", \n", 1<<uint(i), lvname[0])
	}
	fmt.Fprintf(&buff, `
	%v : "%v",
	}
	`, 1<<uint(len(leveldata)), "END")

	for _, lvname := range leveldata {

		fmt.Fprintf(&buff, `
				func %[1]s(format string, v ...interface{}) {
					if !GlobalLogger.IsLevel(LL_%[1]s) {
						return
					}
					s := GlobalLogger.Format2Bytes(1, LL_%[1]s, format, v...)
					err := GlobalLogger.Output(LL_%[1]s,s)
					if err != nil {
						fmt.Println(err)
					}
				}
				`, lvname[0],
		)

		fmt.Fprintf(&buff, `
				func (l *LogBase) %[1]s(format string, v ...interface{}) {
					if !l.IsLevel(LL_%[1]s) {
						return
					}
					s := l.Format2Bytes(1, LL_%[1]s, format, v...)
					err := l.Output(LL_%[1]s,s)
					if err != nil {
						fmt.Println(err)
					}
				}
				`, lvname[0],
		)

	}

	return &buff, nil
}
