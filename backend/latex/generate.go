package latex

import (
	"os"
	"os/exec"
	"text/template"
)

// pdflatex -output-directory=pdf_resume -interaction=nonstopmode latex\templates\resume.tex

const pathToLatex = "latex"

type Resume interface {
	TemplateName() string
}

var resumeTMPL *template.Template

func GeneratePDF(fileName string, outputDir string, resume Resume) error {
	err := generateTEX(fileName, resume)
	if err != nil {
		return err
	}
	cmd := exec.Command(`pdflatex`,
		`-interaction=nonstopmode`,
		`-output-directory=`+outputDir,
		pathToLatex+"/tex_files/"+fileName)
	err = cmd.Run()
	return err
}

func InitTampltes() {
	resumeTMPL = template.Must(template.ParseFiles(pathToLatex +
		"/templates/" + resumeClassicName))
}

func generateTEX(fileName string, resume Resume) error {
	file, err := os.Create(pathToLatex + "/tex_files/" + fileName + ".tex")
	if err != nil {
		return err
	}
	defer file.Close()
	err = resumeTMPL.ExecuteTemplate(file, resume.TemplateName(), resume)
	if err != nil {
		return err
	}
	return nil
}
