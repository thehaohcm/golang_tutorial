package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type friendConnection struct {
	Friends []string `json:"friends"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Tinh02468"
	dbName := "golang_assignment"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func showAllUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDb, err := db.Query("SELECT email, first_name, last_name from user")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var res []user
	for selDb.Next() {
		var user user
		err := selDb.Scan(&user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, user)
	}

	// w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Println("res: ", res)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}

}

func show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConn()
	name := r.URL.Query().Get("name")
	result, err := db.Query("SELECT email, first_name, last_name FROM user WHERE first_name = ?", name)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user user
	if result.Next() {
		err := result.Scan(&user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Post: ", user)
	json.NewEncoder(w).Encode(user)
}

func addFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		resulfFlag := false

		var friendConn friendConnection

		err := json.NewDecoder(r.Body).Decode(&friendConn)
		if err != nil {
			panic(err)
		}
		fmt.Println("body: ", friendConn)
		if len(friendConn.Friends) == 2 {
			//check validation of email
			//implement later

			//check if this value exists into DB already or not
			//implement later
			db := dbConn()
			// result,err:=db.Query("SELECT user1_email, user2_email FROM FRIEND_RELATIONSHIP where user1_email=? and user2_email=?")
			// if err != nil {
			// 	panic(err.Error())
			// }

			//insert database
			insForm, err := db.Prepare("INSERT INTO FRIEND_RELATIONSHIP(user1_email,user2_email,blocked) VALUES(?,?,?)")
			if err != nil {
				panic(err)
			}
			insForm.Exec(friendConn.Friends[0], friendConn.Friends[1], 0)
			insForm.Exec(friendConn.Friends[1], friendConn.Friends[0], 0)
			resulfFlag = true
			defer db.Close()
		} else {
			fmt.Println("Error: Invalided data")
		}
		json.NewEncoder(w).Encode(resulfFlag)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func showListFriend(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		emails, ok := r.URL.Query()["email"]

		if !ok || len(emails[0]) < 1 {
			fmt.Println("Url param 'email' is missing")
			return
		}

		email := emails[0]

		db := dbConn()
		friendRecords, err := db.Query("SELECT user2_email from FRIEND_RELATIONSHIP WHERE user1_email='" + email + "'")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var friends friendConnection

		for friendRecords.Next() {
			var email string
			err := friendRecords.Scan(&email)
			if err != nil {
				panic(err)
			}
			friends.Friends = append(friends.Friends, email)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Println("res: ", friends.Friends)
		if err := json.NewEncoder(w).Encode(friends.Friends); err != nil {
			panic(err)
		}
	}
}

func main() {
	log.Println("Server st arted on: http://localhost:8080")

	http.HandleFunc("/", showAllUser)
	http.HandleFunc("/show", show)
	http.HandleFunc("/addFriend", addFriend)
	http.HandleFunc("/listFriend", showListFriend)

	http.ListenAndServe(":8080", nil)
}
