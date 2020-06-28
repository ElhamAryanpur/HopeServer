package hopeserver

import (
	//db "HopeServer/lib/db"
	js_engine "HopeServer/lib/js_engine"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/gofiber/fiber"
)

// Init is used to define service routes
func Init(app *fiber.App) {
	js := js_engine.New(app)
	compere, pkgDirectory := loadPackages()

	if len(compere) != 0 {
		for _, service := range compere {
			fpd := path.Join(pkgDirectory, service["folder"])
			fp := path.Join(fpd, service["main"])
			dbp := path.Join(fpd, "database.db")
			f, _ := ioutil.ReadFile(fp)
			go js.Init(string(f), dbp)
		}
	}
}

// loadPackages is used to load packages from the directory and do auth
func loadPackages() ([]map[string]string, string) {
	var compere []map[string]string
	var compereFile []byte

	cp, _ := os.Getwd()
	cp = path.Join(cp, "packages")
	exist, _ := exists(cp)
	if exist == false {
		os.MkdirAll(cp, os.ModePerm)
		cpfn := path.Join(cp, "compere.json")
		ioutil.WriteFile(cpfn, []byte("[]"), 0644)
		compereFile = []byte("[]")
	} else {
		cpf := path.Join(cp, "compere.json")
		_, err := exists(cpf)
		if err != nil {
			compereFile = []byte("[]")
		} else {
			compereFile, _ = ioutil.ReadFile(cpf)
		}
	}
	json.Unmarshal(compereFile, &compere)

	return compere, cp
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// TODO: Do Wren Engine and then Lua Engine
