package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID       int
	AuthorID int
	Title    string
	TitleArr []string
}

func RunServer() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(http.ErrServerClosed, err)
	}
}

func dbaction() ([]Book, error) {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
 
		err = rows.Scan(&book.ID, &book.AuthorID, &book.Title)
		if err != nil {
			log.Fatal(err)
		}
		book.TitleArr = strings.Split(book.Title, "")
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
	}

	files := []string{
		"./html/index.html",
	}
	templates, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	books, err := dbaction()
	if err != nil {
		log.Print(err.Error())
		fmt.Println(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	// for i, b := range books {
	// 	fmt.Println(i, b)
	// }

	err = templates.ExecuteTemplate(writer, "base", books)
	if err != nil {
		log.Print(err.Error())
		fmt.Println(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
}

func main() {
	RunServer()
}
