package main

import (
	"io"
	"os"
)

func main() {
	src, _ := os.Open(os.Args[1])
	srcinfo, _ := os.Stat(os.Args[1])
	des, _ := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, srcinfo.Mode())
	io.Copy(des, src)
	src.Close()
	des.Close()
}
