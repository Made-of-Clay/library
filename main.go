package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/index.html"))

	err := tmpl.ExecuteTemplate(w, "shell.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func detailHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /detail")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/detail.html"))

	err := tmpl.ExecuteTemplate(w, "shell.html", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/detail", detailHandler)
	http.HandleFunc("/", indexHandler)

	// Write a 404 handler for all unmatched routes
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/favicon.ico")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Return 404 for any non-static and non root or detail routes
	// http.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//     http.Error(w, "Not Found", http.StatusNotFound)
	// })

	// TODO: Implement logging middleware to log all incoming requests, their methods, and their URLs
	log.SetFlags(log.Lshortfile | log.LstdFlags) // Log with file and line number

	// TODO: Implement authentication middleware to ensure only authorized clients can access the server's endpoints and resources

	// TODO: Implement caching middleware to store and serve static files from a cache for faster response times

	fmt.Println("Starting server on localhost:8008 ...")
	log.Fatal(http.ListenAndServe(":8008", nil))
}
