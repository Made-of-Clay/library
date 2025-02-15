package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println("---- title" + title)
	p, err := loadPage(title)
	if err != nil {
		fmt.Println("OH CRAP OH CRAP OH CRAP >>> ", err)
		fmt.Fprintf(w, "ERROR: %s", err.Error())
		return
	}

	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)

	// fmt.Fprintf(w, `<h1>Editing %s</h1>
	// <form action="/save/" method="POST"
	// 	<textarea name=body>%s</textarea>
	// 	<button>Save</button>
	// </form?`,
	// 	p.Title, p.Body)
}

func indexHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/index.html"))

	foo := &Page{Title: "bar "}
	err := tmpl.ExecuteTemplate(w, "shell.html", foo)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func detailHandler(w http.ResponseWriter, _r *http.Request) {
	fmt.Println("Serving route /detail")

	tmpl := template.Must(template.ParseFiles("templates/shell.html", "templates/detail.html"))

	foo := &Page{Title: "bar "}
	err := tmpl.ExecuteTemplate(w, "shell.html", foo)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// TODO resume this after I better understand template parsing
	// http.HandleFunc("/detail", detailHandler)
	// http.HandleFunc("/", indexHandler)

	// Tutorial
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	// http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	fmt.Println("Starting server on localhost:8008 ...")
	log.Fatal(http.ListenAndServe(":8008", nil))
}
