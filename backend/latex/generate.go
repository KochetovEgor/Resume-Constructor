package latex

import "os/exec"

// pdflatex -output-directory=pdf_resume -interaction=nonstopmode latex\templates\resume.tex

func GeneratePDF(inputDir string, outputDir string) (string, error) {
	cmd := exec.Command(`pdflatex`,
		`-interaction=nonstopmode`,
		`-output-directory=`+outputDir,
		inputDir)

	err := cmd.Run()
	return `resume.pdf`, err
}
