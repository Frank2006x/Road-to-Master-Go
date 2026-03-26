package main

import (
	"log"

	"github.com/Frank2006x/Fibre/src"
)

func main() {
	app := src.SetupApp()
	port:=":3000"
	log.Printf("Server is running %s",port)
	app.Listen(port)
}