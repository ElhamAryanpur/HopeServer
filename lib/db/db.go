package hopeserver

import (
	badger "github.com/dgraph-io/badger/v2"
	hp "HopeServer/lib/helper"
)

// DB is the database class
type DB struct {
	database *badger.DB
	txn      *badger.Txn
}

// a function struct for callbacks
type fn func(string)

// Set the data on DB
func (d *DB) Set(key string, value string) {
	err := d.txn.Set([]byte(hp.Encrypt(key)), []byte(hp.Encrypt(value)))
	handle(err, "DB Set")
}

// Get the data from DB
func (d *DB) Get(key string, callback fn) {
	err := d.database.Update(func(txn *badger.Txn) error {
		res, err := d.txn.Get([]byte(hp.Encrypt(key)))
		handle(err, "Get Low")

		res.Value(func(val []byte) error {
			callback(hp.Decrypt(string(val)))
			return nil
		})
		return nil
	})

	handle(err, "Get Main")
}

// New is constructor
func New(path string) DB {
	db, err := badger.Open(badger.DefaultOptions("data.db"))
	handle(err, "init")
	defer db.Close()

	txn := db.NewTransaction(true)

	d := DB{database: db, txn: txn}
	return d
}

func handle(err error, part string) {
	if err != nil {
		println(part, "DB ERROR:", err)
	}
}
