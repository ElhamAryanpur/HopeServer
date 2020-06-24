package hopeserver

import (
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/robertkrimen/otto"
)

// JS is used for compiling JS code
type JS struct {
	app *fiber.App
	vm  *otto.Otto
}

// GET is used for routing GET requests
func (js *JS) GET(call otto.FunctionCall) otto.Value {

	route := call.Argument(0).String()
	response := call.Argument(1).String()
	varType := call.Argument(2).String()

	js.app.Get(route, func(c *fiber.Ctx) {
		if varType == "string" {
			c.Send(response)
		} else if varType == "variable" {
			if va, err := js.vm.Get(response); err == nil {
				if v, err := va.ToString(); err == nil {
					c.Send(v)
				}
			}
		} else if varType == "function" {
			va, _ := js.vm.Run(response + "()")
			val, _ := va.ToString()
			c.Send(val)
		}
	})
	return otto.Value{}
}

// Init is used to Initialize stuff
func (js *JS) Init() {
	js.vm.Set("GET", js.GET)

	f, _ := ioutil.ReadFile("packages/test/main.js")
	js.vm.Run(string(f))
}

// New is used to initialize lua stuff
func New(app *fiber.App) JS {
	vm := otto.New()
	return JS{app: app, vm: vm}
}
