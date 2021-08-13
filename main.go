package main

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

type BonUrl struct {
	Url string `json:"url" xml:"url" form:"url"`
}

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
		b := new(BonUrl)

		if err := c.BodyParser(b); err != nil {
			fmt.Println(err)
			return c.SendStatus(500)
		}
		idRaw := make([]byte, 2)
		rand.Read(idRaw)

		sid := hex.EncodeToString(idRaw)
		slink := c.BaseURL() + "/" + sid

		AddLink(sid, b.Url)
		return c.JSON(NewURL{slink, sid})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		sid := c.Params("id")
		target := GetLink(sid)

		return c.Redirect(target)
	})

	app.Listen(":3000")
}

