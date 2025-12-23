package main

import (
	"encoding/json"
	"fmt"
	"os"
	"resume-backend/latex"
)

const pathToLatex = `/latex`

func testJSONEncoder() error {
	file, err := os.Open("test_jsons/test1.json")
	if err != nil {
		return err
	}
	defer file.Close()
	var resume latex.ResumeClassic
	err = json.NewDecoder(file).Decode(&resume)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(&resume, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	return nil
}

func main() {
	fmt.Println(latex.GeneratePDF(`latex\templates\resume_friend.tex`, `pdf_resume`))
	//fmt.Println(testJSONEncoder())
}
