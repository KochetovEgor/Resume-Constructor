package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
}

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/index.html")
}

func (a *App) Run() error {
	http.HandleFunc("/", a.mainPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server := &http.Server{
		Addr:              ":5100",
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Printf("Frontend server started on: %v\n", server.Addr)
	err := server.ListenAndServe()
	file, _ := os.Open("ttt.txt")
	fmt.Println(file)
	return err
}
