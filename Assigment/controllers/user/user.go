package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	userResponse "Assigment/models/user"
	ultis "Assigment/ultilities"
)

// ShowAllUser godoc
// @Summary show all infos of all users
// @Description return a list of entire users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /showAllUsers [get]
func ShowAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := ultis.DBConn()
		selDb, err := db.Query("SELECT email, first_name, last_name from user")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var res []userResponse.User
		for selDb.Next() {
			var user userResponse.User
			var firstName, lastName sql.NullString
			err := selDb.Scan(&user.Email, &firstName, &lastName)
			if err != nil {
				panic(err.Error())
			}
			if firstName.Valid {
				user.FirstName = firstName.String
			}
			if lastName.Valid {
				user.LastName = lastName.String
			}
			res = append(res, user)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("res: ", res)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
