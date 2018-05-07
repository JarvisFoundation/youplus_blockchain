package main

import (
	"runtime"

	"./app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var app = new(app.App)
	app.Init()
}
