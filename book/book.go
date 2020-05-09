package book

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	"go_fiber/database"
)

// Book struct
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks function
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetSingleBook function
func GetSingleBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

//AddBook function
func AddBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		log.Fatal(err)
	}
	db.Create(&book)
	c.JSON(book)
}

//DeleteBook function
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(404).Send("no book found")
		return
	}
	db.Delete(&book)
	c.Send("Delete book successfully")
}
