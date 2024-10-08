package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	// Add handler to display "pong" in console
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("pong")
		w.Write([]byte("pong"))
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
