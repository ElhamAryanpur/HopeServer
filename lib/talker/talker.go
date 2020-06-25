package hopeserver

import (
	//db "HopeServer/lib/db"
	js_engine "HopeServer/lib/js_engine"
	"io/ioutil"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	js := js_engine.New(app)

	f, _ := ioutil.ReadFile("packages/test/main.js")
	go js.Init(string(f))
}

// TODO: Do Wren Engine and then Lua Engine
