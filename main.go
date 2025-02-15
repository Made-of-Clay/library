package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func indexHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/index.html"))

	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func detailHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /detail")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/detail.html"))

	err := tmpl.ExecuteTemplate(w, "detail.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/detail", detailHandler)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
