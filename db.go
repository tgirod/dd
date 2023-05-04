package main

import (
	"fmt"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/codec/gob"
	"github.com/asdine/storm/v3/q"
)

var db = func() *storm.DB {
	db, err := storm.Open("dirtydistrict.db", storm.Codec(gob.Codec))
	if err != nil {
		panic(err)
	}
	return db
}()

// One récupère un enregistrement à partir d'un champ
func One[T any](field string, value any) (T, error) {
	var one T
	if err := db.One(field, value, &one); err != nil {
		return one, fmt.Errorf("%s=%v : %w", field, value, err)
	}
	return one, nil
}

// Delete supprime un enregistrement
func Delete[T any](value T) error {
	return db.DeleteStruct(&value)
}

// Save sauve un enregistrement
// met à jour un existant ou créé un nouveau selon la valeur du champ id
func Save[T any](value T) (T, error) {
	err := db.Save(&value)
	return value, err
}

// Update met à jour un enregistrement déjà existant
func Update[T any](value T) (T, error) {
	err := db.Update(&value)
	return value, err
}

// Find récupère une liste d'enregistrements correspondant à une requête
func Find[T any](matchers ...q.Matcher) ([]T, error) {
	var values []T
	err := db.Select(matchers...).Find(&values)
	if err == storm.ErrNotFound {
		return values, nil
	}
	return values, err
}

// First récupère le premier enregistrement correspondant à une requête
func First[T any](matchers ...q.Matcher) (T, error) {
	var one T
	err := db.Select(matchers...).First(&one)
	return one, err
}
