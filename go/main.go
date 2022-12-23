/*
The MIT License (MIT)

Copyright (c) 2022 sodamouse

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func getList(filepath string) ([]string, error) {
	data, e := os.ReadFile(filepath)

	if e != nil {
		return nil, e
	}

	return strings.Split(string(data), "\n"), nil
}

func remove(homeDir string, item string) {
	path := homeDir + "/" + item
	stat, _ := os.Stat(path)

	if stat == nil {
		return
	}

	e := os.RemoveAll(path)

	if e == nil {
		fmt.Println("Removed:", item)
	} else {
		fmt.Println("Could not remove: ", item)
	}
}

func main() {
	programVersion := "clean 1.1 (Go)"
	userHome, e := os.UserHomeDir()
	checkFatal(e)

	var listFilePath string
	flag.StringVar(&listFilePath, "l", userHome+"/.config/clean/files.list", "Specifies the path to list file")
	showVersion := flag.Bool("v", false, "Displays program version information")
	flag.Parse()

	if *showVersion {
		fmt.Println(programVersion)
		return
	}

	list, e := getList(listFilePath)
	checkFatal(e)

	for _, item := range list {
		if item == "" {
			continue
		}

		remove(userHome, item)
	}
}

func checkFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
