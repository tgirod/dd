package main

import (
	"flag"
)

const host = "0.0.0.0"
const portPlayer = 1337
const portMonitor = 7331

var app *App

func main() {
	init := flag.Bool("init", false, "réinitialise la base de données")
	flag.Parse()
	app = NewApp(*init)
	app.Start()
}
