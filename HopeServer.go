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

	PublicFolder, _ := os.Getwd()
	PublicFolder = path.Join(PublicFolder, "front", "public")

	app.Get("/", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, "index.html"))
	})

	app.Get("/public/:file", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, c.Params("file")))
	})

	app.Get("/public/build/:buildfile", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, "build", c.Params("buildfile")))
	})

	// Route The Public Folder

	app.Listen(3000)
}
