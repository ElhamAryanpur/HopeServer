package hopeserver

import (
	//db "HopeServer/lib/db"
	js_engine "HopeServer/lib/js_engine"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	js := js_engine.New(app)
	js.Init()

	app.Get("/service/register", func(c *fiber.Ctx) {
		c.Send("AAAAA")
	})
}

// TODO: Do Wren Engine and then Lua Engine
