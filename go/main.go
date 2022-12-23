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
	userHome, e := os.UserHomeDir()
	checkFatal(e)

	var listFilePath string
	flag.StringVar(&listFilePath, "l", userHome+"/.config/clean/files.list", "Specifies the path to list file")
	flag.Parse()

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
