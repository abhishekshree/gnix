package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(path.Dir(os.Args[1]))
	}
}
