package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	os.Mkdir(flag.Arg(0), 0644)
}
