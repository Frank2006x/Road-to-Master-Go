package main

import (
	"urlshorter/src"
)

func main() {
	app := src.SetupApp()
	

	app.Listen(":3000")
}