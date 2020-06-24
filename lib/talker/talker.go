package hopeserver

import (
	//db "HopeServer/lib/db"
	js_engine "HopeServer/lib/js_engine"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	js_engine.Init(*app)

	app.Get("/service/register", func(c *fiber.Ctx) {
		c.Send("AAAAA")
	})
}
