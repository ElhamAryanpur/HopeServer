package main

import (
	tk "HopeServer/lib/talker"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	tk.Init(app)

	app.Get("/nosql/set/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}
