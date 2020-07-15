package hopeserver

import (
	db "HopeServer/lib/db"

	"github.com/dgraph-io/badger/v2"
	"github.com/gofiber/fiber"
	"github.com/robertkrimen/otto"
)

// JS is used for compiling JS code
type JS struct {
	DB  *db.DB
	app *fiber.App
	vm  *otto.Otto
}

// NoSQLGet is used for routing HTTP GET requests
func (js *JS) NoSQLGet(call otto.FunctionCall) otto.Value {

	key := call.Argument(0).String()
	callback := call.Argument(1).String()

	js.DB.Get(key, func(result string) {
		res, _ := js.vm.ToValue(result)
		r, _ := res.ToString()
		js.vm.Run(callback + "('" + r + "');")
	})

	return otto.Value{}
}

// NoSQLSet is used for routing HTTP GET requests
func (js *JS) NoSQLSet(call otto.FunctionCall) otto.Value {

	key := call.Argument(0).String()
	value := call.Argument(1).String()
	js.DB.Set(key, value)

	return otto.Value{}
}

// HTTPGet is used for routing HTTP GET requests
func (js *JS) HTTPGet(call otto.FunctionCall) otto.Value {

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

// InitNoSQL is used to Initialize the database
func (js *JS) InitNoSQL(path string) {
	ddb, err := badger.Open(badger.DefaultOptions(path))
	println(err)
	defer ddb.Close()
	d := db.New(path)
	d.SetDB(ddb)
	js.DB = &d
}

// Init is used to Initialize stuff
func (js *JS) Init(data string, dbPath string) {
	js.InitNoSQL(dbPath)
	js.vm.Set("HTTPGet", js.HTTPGet)
	js.vm.Set("NoSQLSet", js.NoSQLSet)
	js.vm.Set("NoSQLGet", js.NoSQLGet)
	js.vm.Run(data)
}

// New is used to initialize JS stuff
func New(app *fiber.App) JS {
	vm := otto.New()
	return JS{app: app, vm: vm}
}
