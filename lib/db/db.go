package hopeserver

import (
	"io/ioutil"
	hp "HopeServer/lib/helper"
)

// Set the data on DB
func Set(DB *bitcask.Bitcask, key string, value string) {
	
}

// Get the data from DB
func Get(DB *bitcask.Bitcask, key string) {
	
}

// Init is constructor
func Init(path string) {
	db, _ := ioutil.ReadFile(path)
}
