package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"resume-backend/latex"
	"time"
)

const frontendServerAddr = "http://localhost:5100"

const pdfDir = "pdf_resume"

const basePDFname = "resume"

type App struct {
}

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user!!!")
}

// Return user's resume in pdf format
func (a *App) getResume(w http.ResponseWriter, r *http.Request) {
	resume := &latex.ResumeClassic{}
	err := json.NewDecoder(r.Body).Decode(resume)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error decoding test JSON: %v\n", err)
		return
	}

	dir, err := os.MkdirTemp(pdfDir, "pdf_")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating temp dir: %v\n", err)
	}
	defer os.RemoveAll(dir)

	err = latex.GeneratePDF(basePDFname, dir, resume)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error generating pdf: %v\n", err)
		return
	}

	filePath := filepath.Join(dir, basePDFname+".pdf")
	file, err := os.Open(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error opening pdf file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Printf("error copying pdf file to client: %v\n", err)
		return
	}
}

func (a *App) debugPage(w http.ResponseWriter, r *http.Request) {
	resume := &latex.ResumeClassic{}
	err := json.NewDecoder(r.Body).Decode(resume)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error decoding test JSON: %v\n", err)
		return
	}
	data, _ := json.MarshalIndent(resume, " ", "  ")
	fmt.Printf("%s\n", data)
}

// For allowing CORS
func allowCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontendServerAddr)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Starts server
func (a *App) Run() error {
	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("POST /resume", a.getResume)
	http.HandleFunc("/debug", a.debugPage)

	server := &http.Server{
		Addr:              ":8100",
		Handler:           allowCORS(http.DefaultServeMux),
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("Backend server started on: %v\n", server.Addr)
	err := server.ListenAndServe()
	return err
}
