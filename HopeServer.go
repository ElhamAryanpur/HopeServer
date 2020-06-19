package main

import (
	"github.com/gofiber/fiber"

	//"github.com/syndtr/goleveldb/leveldb"
	//tk "HopeServer/lib/talker"
)

func main() {
  app := fiber.New()

  app.Get("/nosql/set/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })

  app.Listen(3000)
}