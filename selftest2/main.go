package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In search of life vol1", Author: "anshu", Quantity: 2},
	{ID: "2", Title: "In search of life vol2", Author: "anshu", Quantity: 3},
	{ID: "2", Title: "In search of life vol3", Author: "anshu", Quantity: 4},
	{ID: "4", Title: "In search of life vol4", Author: "anshu", Quantity: 5},
}

func getAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)

}
func getBookByID(c echo.Context) error {
	id := c.Param("id")
	author := c.Param("author")

	for _, book := range books {
		if book.Author == author {
			if book.ID == id {
				return c.JSON(http.StatusOK, book)
			}
		}
	}

	return c.String(http.StatusNotFound, "Book not found")
}
func addBooks(c echo.Context) error {
	var book Book
	if err := c.Bind(&book); err != nil {
		return err
	}
	book.ID = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	return c.JSON(http.StatusCreated, book)

}
func count(c echo.Context) error {
	id := c.Param("id")
	count := 0
	for _, book := range books {
		if book.ID == id {
			count++
		}
	}
	return c.JSON(http.StatusOK, map[string]int{"count": count})

}
func countByAuthor(c echo.Context) error {
	author := c.Param("author")
	totalQuantity := 0
	for _, book := range books {
		if book.Author == author {
			totalQuantity += book.Quantity

		}
	}
	return c.JSON(http.StatusOK, map[string]int{"count": totalQuantity})

}


func addBook(title, author string, quantity int) {
	books = append(books, Book{
		ID:       strconv.Itoa(len(books) + 1),
		Title:    title,
		Author:   author,
		Quantity: quantity,
	})
}

func main() {
	e := echo.New()
	e.GET("/books", getAllBooks)
	e.GET("/books/:author/:id", getBookByID)
	e.POST("/books", addBooks)
	e.POST("/booksm", addBook)
	e.GET("/count/:id", count)
	e.GET("/count/:author", countByAuthor)
	e.Logger.Fatal(e.Start(":8220"))

}
