package main

import "github.com/asdine/storm/v3"
import "github.com/asdine/storm/v3/codec/gob"

var db = func() *storm.DB {
	db, err := storm.Open("dirtydistrict.db", storm.Codec(gob.Codec))
	if err != nil {
		panic(err)
	}
	return db
}()

// Init réinitialise la BDD avec les données par défaut
func Init() {
	structs := []any{
		Identity{},
		Message{},
		Server{},
		Account{},
		Link{},
		Entry{},
		Register{},
		Post{},
	}
	for _, s := range structs {
		db.Drop(&s)
		db.Init(&s)
	}
}
