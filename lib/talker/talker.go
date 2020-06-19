package hopeserver

import (
	db "HopeServer/lib/db"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	DB := db.Init("data")
	db.Set(DB, "k", "v")
	DB.Put([]byte("Hello"), []byte("World"))
	val, _ := DB.Get([]byte("Hello"))
	println(val)
	println(db.Get(DB, "k"))
	app.Get("/service/register", func(c *fiber.Ctx) {
		c.Send("AAAAA")
	})
}
