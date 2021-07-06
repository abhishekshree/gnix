package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {

	cu, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cu.Username)
}
