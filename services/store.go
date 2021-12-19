package services

import (
	"encoding/json"
	"errors"
	"fmt"
	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/models"
	guuid "github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Store struct {
	Db models.Database
}

func DB(db models.Database) Store {
	store := Store{Db: db}
	ReadFileFromDB(&db)

	defer func() {
		go SetInterval(store)
	}()

	return store
}

/*
 - Set Service Method.
 - Save key and value to Store
 - Returns two values(string, error)
*/
func (store Store) Set(value string) (string, error) {

	if value == "" {
		return "", errors.New(CONSTANTS.SETDATA_ERR_MSG)
	}

	id := guuid.New().String()
	store.Db[id] = value
	return id, nil
}

/*
 - Get Service Method.
 - Get value by key from the Store
 - Returns two values(string, error)
*/
func (store Store) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New(CONSTANTS.GETDATA_ERR_MSG)
	}
	return store.Db[key], nil
}

/*
 - GetAll Service Method.
 - Get all values from the Store
 - Returns two values(string, error)
*/
func (store Store) GetAll() (map[string]string, error) {
	mValues := make(map[string]string)

	for key, value := range store.Db {
		mValues[key] = value
	}

	return mValues, nil
}

/*
 - FlushData Service Method.
 - Delete all values from the Store
 - Returns one value(error)
*/
func (store Store) FlushData() error {
	for i := range store.Db {
		delete(store.Db, i)
	}
	size := len(store.Db)
	if size != 0 {
		return errors.New(CONSTANTS.FLUSHDATA_ERR_MSG)
	}
	return nil
}

/*
 - DeleteSingle Service Method.
 - Delete value by key from the Store
 - Returns one value(error)
*/
func (store Store) DeleteSingle(key string) error {
	if store.Db[key] == "" {
		return errors.New(CONSTANTS.DELETESINGLE_ERR_MSG)
	}
	delete(store.Db, key)
	value := store.Db[key]
	if value != "" {
		return errors.New(CONSTANTS.DELETESINGLE_ERR_MSG)
	}
	return nil
}

/*
 - SetInterval Method
 - Write File every 12 seconds.
*/
func SetInterval(store Store) {
	timer := time.NewTicker(12 * time.Second)
	for _ = range timer.C {
		dbJson, _ := json.Marshal(store.Db)
		fmt.Println("SetIntervalCalled", dbJson)
		ioutil.WriteFile("temp/data.json", dbJson, os.ModeAppend)
	}
}

/*
 - ReadFileFromDB Method
 - ReadFile from "emp/data.json"
*/
func ReadFileFromDB(db *models.Database) {
	bytes, err := ioutil.ReadFile("temp/data.json")
	fmt.Println(bytes)
	if err != nil {
		log.Panic("Err:", err.Error())
	}
	json.Unmarshal(bytes, &db)
}
