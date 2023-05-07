package main

import (
	"flag"
)

const host = "0.0.0.0"
const port = 1337

func main() {
	init := flag.Bool("init", false, "réinitialise la base de données")
	flag.Parse()
	a := NewApp(*init)
	a.Start()
}
