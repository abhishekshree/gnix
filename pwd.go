package main

import (
	"fmt"
	"os"
)

func main() {
	workingdir, _ := os.Getwd()
	fmt.Println(workingdir)
}
