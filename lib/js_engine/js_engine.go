package hopeserver

import (
	"github.com/gofiber/fiber"
	"github.com/robertkrimen/otto"
)

func test(call otto.FunctionCall, router *fiber.App) otto.Value{
	println("YOU SAID: %s", call.Argument(0).String());
	router.Get(call.Argument(0).String(), func (c *fiber.Ctx)  {
		// Make route with d as query
		// Then passs query data to function call given.
	})
	return otto.Value{};
}

// Init is used to initialize lua stuff
func Init(app *fiber.App){
	vm := otto.New()
	vm.Set("test", test);
	vm.Run("test('AAA')");
}