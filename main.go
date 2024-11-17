package main

import (
	"book-catalogue-service/src/books"
	"book-catalogue-service/src/db"
	"book-catalogue-service/src/server"
)

func main() {
	db.InitDB()
	db.Get().AutoMigrate(books.Book{})
	server.Start()
}
