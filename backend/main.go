package main

import (
	"fmt"
	"resume-backend/app"
	"resume-backend/latex"
)

func main() {
	latex.InitTemplates()
	app := &app.App{}
	fmt.Println(app.Run())
}
