package main

import (
	"bookstore-rest-api/internal/service"
	"bookstore-rest-api/internal/store"
	"bookstore-rest-api/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db_book.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	q := `CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	booksStore := store.New(db)
	bookService := service.New(booksStore)
	bookHandler := transport.New(bookService)

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Server is running on http://localhost:8080")
	fmt.Println("API Endpoints:")
	fmt.Println("GET /books - Get all books")
	fmt.Println("POST /books - Create a new book")
	fmt.Println("GET /books/{id} - Get a book by ID")
	fmt.Println("PUT /books/{id} - Update a book by ID")
	fmt.Println("DELETE /books/{id} - Delete a book by ID")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
