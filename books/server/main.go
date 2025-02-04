package main

import (
	"com/gogo/config"
	"com/gogo/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	db := config.InitDB()

	app.Use(cors.New())

	app.Get("/books", func(c fiber.Ctx) error {
		var books []entity.Book 
		if err := db.Find(&books).Error; err != nil {
			return c.Status(500).SendString("Failed to fetch books: " + err.Error())
		}

		return c.JSON(books)
	})

	app.Listen(":3000")
}
