package main

import (
	"github.com/gofiber/fiber"
	tk "HopeServer/lib/talker"
)

func main() {
  app := fiber.New()

  tk.Init(app)

  app.Get("/nosql/set/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })

  app.Listen(3000)
}