package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var dir string
	lenArgs := len(os.Args)
	if lenArgs > 1 {
		dir = os.Args[lenArgs-1]
	} else {
		dir = "."
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		fmt.Println(f.Name())
	}
}
