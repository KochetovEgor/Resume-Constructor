package main

import (
	"frontend/app"
	"log"
)

func main() {
	app := &app.App{}
	log.Println(app.Run())
}
