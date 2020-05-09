package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go_fiber/book"
	"go_fiber/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:@/mahasiswa?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection succesfully opened!")

	// database.DBConn.AutoMigrate(&book.Book{})
	// fmt.Println("Database Migrated!")
}

func CheckMidlleware(c *fiber.Ctx) {
	fmt.Println("Middleware trigered")
	c.Next()
}

func CheckMidlleware2(c *fiber.Ctx) {
	token := c.Get("Authorization")
	fmt.Println(token)
	c.Next()
}

func setUpRoute(app *fiber.App) {
	app.Get("/books", CheckMidlleware, book.GetBooks)
	app.Get("/book/:id", CheckMidlleware2, book.GetSingleBook)
	app.Post("/book", book.AddBook)
	app.Delete("/book/:id", book.DeleteBook)
}

func helloWorld(c *fiber.Ctx) {
	c.Send("hallo world frm func")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setUpRoute(app)

	app.Listen(8080)
}
