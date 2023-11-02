package main

import (
	"html/template"
	"log"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", handlePage)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(http.ErrServerClosed, err)
	}
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
	}

	files := []string{
		"./html/base.tmpl",
		"./html/index.tmpl",
	}
	templates, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
}

func main() {
	RunServer()
}



