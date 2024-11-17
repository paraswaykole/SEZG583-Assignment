package server

import (
	"book-catalogue-service/src/books"
	"book-catalogue-service/src/constants"
	"book-catalogue-service/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"health": "ok",
		})
	})

	api := app.Group("/api")
	api.Use(middleware.JWTMiddleware)
	booksGroup := api.Group("/books")

	booksGroup.Post("/create", books.CreateBookController)
	booksGroup.Get("/list", books.GetBooksListController)
	booksGroup.Get("/:id", books.GetBookByIdController)
	booksGroup.Delete("/:id", books.DeleteBookController)

	app.Listen(constants.SERVER_PORT)
}
