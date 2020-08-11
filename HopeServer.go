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

	app.Get("/:file", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, c.Params("file")))
	})

	app.Get("/build/:buildfile", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, "build", c.Params("buildfile")))
	})

	app.Get("/themes/:themefile", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, "themes", c.Params("themefile")))
	})

	app.Get("/meta/:metafile", func(c *fiber.Ctx) {
		c.SendFile(path.Join(PublicFolder, "meta", c.Params("metafile")))
	})

	app.Listen(3000)
}
