package main

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Bon",
		})
	})

	app.Post("/api/new", func(c *fiber.Ctx) error {
		url := c.Query("url")
		fmt.Println(url)

		idRaw := make([]byte, 2)
		rand.Read(idRaw)

		urlid := hex.EncodeToString(idRaw)
		shortlink := c.BaseURL() + "/" + urlid

		return c.JSON(NewURL{shortlink, urlid})
	})

	app.Listen(":3000")
}

