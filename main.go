package main

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	InitDB()
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine.Reload(true),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Home",
		})
	})

	app.Post("/api/new", func(c *fiber.Ctx) error {
		url := c.Query("url")
		fmt.Println(url)

		idRaw := make([]byte, 2)
		rand.Read(idRaw)

		sid := hex.EncodeToString(idRaw)
		slink := c.BaseURL() + "/" + sid

		AddLink(sid, url)
		return c.JSON(NewURL{slink, sid})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		sid := c.Params("id")
		target := GetLink(sid)

		return c.Redirect(target)
	})

	app.Listen(":3000")
}

