package main

import (
	"io"
	"log"
	"os"
)

func catFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	return err
}

func main() {
	for i := 1; i+1 <= len(os.Args); i++ {
		if err := catFile(os.Args[i]); err != nil {
			log.Fatal(err)
		}
	}
}
