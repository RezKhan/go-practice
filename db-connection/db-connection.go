package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
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

	for rows.Next() {
		var id int
		var author_id int
		var title string

		err = rows.Scan(&id, &author_id, &title)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, author_id, title)
	}
}
