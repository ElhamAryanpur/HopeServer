package main

import (
	ft "HopeServer/lib/front"
	tk "HopeServer/lib/talker"
	"os"
	"path"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	ft.Init(app)
	tk.Init(app)

	app.Get("/", func(c *fiber.Ctx) {
		indexFile, _ := os.Getwd()
		indexFile = path.Join(indexFile, "front", "public", "index.html")
		c.SendFile(indexFile)
	})

	// Route The Public Folder

	app.Listen(3000)
}
