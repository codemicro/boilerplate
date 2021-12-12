package db

import (
	bh "github.com/timshannon/bolthold"
)

type DB struct {
	store *bh.Store
}

func New(filename string) (*DB, error) {
	store, err := bh.Open(filename, 0644, nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		store: store,
	}, nil
}