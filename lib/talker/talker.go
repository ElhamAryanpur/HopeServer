package hopeserver

import (
	db "HopeServer/lib/db"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	DB := db.New("data.db")
	DB.Set("key", "value")
	DB.Get("key", func (result string)  {
		println(result)
	})

	app.Get("/service/register", func(c *fiber.Ctx) {
		c.Send("AAAAA")
	})
}
