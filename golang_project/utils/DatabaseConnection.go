package utils

import (
	"database/sql"
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

var db *sql.DB

//singleton pattern
func GetInstance() *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance now.")
			db = connectDatabase()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return db
}

func connectDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "C:/Users/hao.nguyen/Desktop/golang_tutorial/golang_project/golang_project.db")
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to db")
	return db
}
