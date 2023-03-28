package main

import (
	"fmt"
	"os"
)

const host = "0.0.0.0"
const port = 1337

func main() {
	// Command line filename to init game with
	args := os.Args
	fmt.Printf("args = %v \n", args)
	fmt.Printf("len(args)= %d \n", len(args))
	a := NewApp()
	if len(args) > 1 {
		a.Start(args[1])
	} else {
		a.Start("")
	}

}
