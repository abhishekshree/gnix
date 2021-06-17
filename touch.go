package main

import (
	"os"
	"time"
)

func main() {
	fname := os.Args[1]
	if _, err := os.Stat(fname); err == nil {
		now := time.Now()
		os.Chtimes(fname, now, now)
	} else {
		f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0644)
		f.Close()
	}
}
