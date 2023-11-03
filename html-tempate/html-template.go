package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type BirthdayDetails struct {
	Name string
	Month string
	Day string
}


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
		"./html/base.html",
		"./html/index.html",
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

	birthdays := []BirthdayDetails{} 

	if request.Method == http.MethodPost {
		birthday := BirthdayDetails {
			Name: request.FormValue("name"),
			Month: request.FormValue("month"),
			Day: request.FormValue("day"),
		}
		birthdays = append(birthdays, birthday)
	}

	fmt.Println(birthdays)
}

func main() {
	RunServer()
}



