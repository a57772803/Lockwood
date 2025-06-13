package db_test

import (
	"encoding/json"
	"fmt"
	"lockwood/models"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func getDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestGetDB(t *testing.T) {
	db, err := getDB("testdb")

	fmt.Println(db)
	fmt.Print(err)
}

func TestDbInsert(t *testing.T) {
	db, err := getDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.Put([]byte("test"), []byte("123"), nil)
	db.Put([]byte("test"), []byte("444"), nil)

	data, err := db.Get([]byte("test"), nil)
	fmt.Println(data)
	fmt.Print(err)

	err = db.Delete([]byte("test"), nil)
	fmt.Println(data)
	fmt.Print(err)
}

func TestDbInsertPassword(t *testing.T) {
	db, err := getDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	password := models.PasswordEntry{
		ID:       "test",
		Account:  "test",
		Password: "123",
		Note:     "test",
	}

	jsonData, err := json.Marshal(password)
	if err != nil {
		t.Fatal(err)
	}

	//db.Put([]byte(password.ID), []byte(password.Password), nil)
	db.Put([]byte(password.ID), jsonData, nil)

	data, err := db.Get([]byte(password.ID), nil)
	fmt.Println(string(data))
	fmt.Print(err)

	err = db.Delete([]byte("test"), nil)
	fmt.Print(err)
}

func TestGetAllRecord(t *testing.T) {
	db, err := getDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	_ = db.Put([]byte("user:1"), []byte("leon"), nil)
	_ = db.Put([]byte("user:2"), []byte("gpt"), nil)
	_ = db.Put([]byte("config:theme"), []byte("dark"), nil)

	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	fmt.Printf("----- Dumping LevelDB -----\n")
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		fmt.Printf("Key: %s | Value: %s\n", key, value)
	}
}
