package hopeserver

import (
	tk "HopeServer/lib/talker"
	"encoding/json"

	"github.com/gofiber/fiber"
)

// Init all the needed routes for front
func Init(app *fiber.App) {

	app.Get("/meta/services", func(c *fiber.Ctx) {

		services, _ := tk.LoadPackages()
		var list []string

		for _, s := range services {
			list = append(list, s["name"])
		}

		result, _ := json.Marshal(list)
		c.Send(result)
	})

}
