package main

import (
	"encoding/json"
	"fmt"
	"os"
	"resume-backend/latex"
)

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

func testJSONDecoder(resume latex.Resume) error {
	file, err := os.Open("test_jsons/test1.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(resume)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	latex.InitTampltes()
	resume := &latex.ResumeClassic{}
	fmt.Println(testJSONDecoder(resume))
	fmt.Println(latex.GeneratePDF("test", "pdf_resume", resume))
}
