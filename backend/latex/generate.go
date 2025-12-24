package latex

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// Path to this package from main.go
const pathToLatex = "latex"

var templateNames = [...]string{
	resumeClassicName,
}

// Type for all resume templates
type Resume interface {
	TemplateName() string
}

// Variable for all templates
var resumeTMPL *template.Template

// Generates resume in .pdf format in directory outputDir. Variable resume is user's data.
func GeneratePDF(fileName string, outputDir string, resume Resume) error {
	file, err := os.Create(filepath.Join(outputDir, fileName+".tex"))
	if err != nil {
		return fmt.Errorf("error creating .tex file: %v\n", err)
	}

	err = generateTEX(file, resume)
	file.Close()
	if err != nil {
		return fmt.Errorf("error generating .tex file: %v\n", err)
	}

	cmd := exec.Command(`pdflatex`,
		`-interaction=nonstopmode`,
		`-output-directory=`+outputDir,
		file.Name())
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating pdf: %v\n", err)
	}
	return nil
}

// Initializes all resume templates
func InitTemplates() {
	templatePaths := make([]string, len(templateNames))
	for i, tmpl := range templateNames {
		templatePaths[i] = filepath.Join(pathToLatex, "templates", tmpl)
	}
	resumeTMPL = template.Must(template.ParseFiles(templatePaths...))
}

// Generates user's resume, but in .tex format, which should then be converted to .pdf format
func generateTEX(wr io.Writer, resume Resume) error {
	err := resumeTMPL.ExecuteTemplate(wr, resume.TemplateName(), resume)
	if err != nil {
		return fmt.Errorf("error executing template %v: %v\n", resume.TemplateName(), err)
	}
	return nil
}
