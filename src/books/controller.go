package books

import (
	"github.com/gofiber/fiber/v2"
)

func CreateBookController(c *fiber.Ctx) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(map[string]interface{}{
			"message": "incorrect book format",
		})
	}
	resBook, err := CreateBook(&book)
	if err != nil {
		panic(err)
	}
	return c.JSON(resBook)
}

func GetBooksListController(c *fiber.Ctx) error {
	books, err := GetBooks(0, 100, "")
	if err != nil {
		panic(err)
	}
	return c.JSON(books)
}

func GetBookByIdController(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)
	if err != nil {
		panic(err)
	}
	if id == -1 {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]interface{}{
			"message": "id not specified",
		})
	}
	book, err := GetBookByID(uint(id))
	if err != nil {
		panic(err)
	}
	return c.JSON(book)
}

func DeleteBookController(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)
	if err != nil {
		panic(err)
	}
	if id == -1 {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]interface{}{
			"message": "id not specified",
		})
	}
	err = DeleteBook(uint(id))
	if err != nil {
		panic(err)
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"message": "Successfully delete book",
	})
}
