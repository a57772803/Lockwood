package db

import (
	"encoding/json"
	"lockwood/models"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func GetDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *leveldb.DB) error {
	err := db.Close()
	return err
}

func SavePassword(entry models.PasswordEntry) error {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return db.Put([]byte(entry.ID), jsonData, nil)
}

func GetPassword(id string) (*models.PasswordEntry, error) {
	data, err := db.Get([]byte(id), nil)
	if err != nil {
		return nil, err
	}
	var entry models.PasswordEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, err
	}
	return &entry, nil
}
