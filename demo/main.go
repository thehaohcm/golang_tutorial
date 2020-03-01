package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type BankAccount struct {
	id        int
	full_name string
	balance   int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Aa123456@"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nill {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDb, err := db.Query("SELECT * from new_schema.bank_account order by id desc")
	if err != nil {
		panic(err.Error())
	}
	bankAccount := BankAccount{}
	res := []BankAccount{}
	for selDb.Next() {
		var id int
		var fullName string
		var balance int

		err = selDB.Scan(&id, &fullName, &balance)
		if err != nil {
			panic(err.Error())
		}
		bankAccount.id = id
		bankAccount.full_name = fullName
		bankAccount.balance = balance
		res = append(res, bankAccount)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * from new_schema.bank_account where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	bankAccount := BankAccount{}
	for selDb.Next() {
		var id int
		var fullname string
		var balance int
		err = selDB.Scan(&id, &fullname, &balance)
		if err != nil {
			panic(err.Error())
		}
		bankAccount.id = id
		bankAccount.full_name = fullname
		bankAccount.balance = balance
	}
	tmpl.ExecuteTemplate(w, "Show", bankAccount)
	defer db.Close()
}

func main() {
	log.Println("Server st arted on: http://localhost:8080")

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)

	http.ListenAndServe(":8080", nil)
}
