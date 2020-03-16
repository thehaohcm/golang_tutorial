package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	response "Assigment/models/commons"
	friendResponse "Assigment/models/friend"
	ultis "Assigment/ultilities"
)

// // show godoc
// // @Summary add a new friend
// // @Description return a result of creating a new friend by 2 user emails
// // @Tags Friend
// // @Accept  json
// // @Produce  json
// // @Param friend body models.FriendRequest true "friend"
// // @Success 200 {object} models.Response
// // @Router /addFriend [post]
// func show(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		db := dbConn()
// 		name := r.URL.Query().Get("name")
// 		result, err := db.Query("SELECT email, first_name, last_name FROM user WHERE first_name = ?", name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		defer result.Close()
// 		var user user
// 		if result.Next() {
// 			err := result.Scan(&user.Email, &user.FirstName, &user.LastName)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 		}
// 		fmt.Println("Post: ", user)
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 	}
// }

// AddFriend godoc
// @Summary add a new friend
// @Description return a result of creating a new friend by 2 user emails
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param friend body models.FriendRequest true "friend"
// @Success 200 {object} models.Response
// @Router /addFriend [post]
func AddFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		resultFlag := false

		var friendRes friendResponse.FriendRequest

		err := json.NewDecoder(r.Body).Decode(&friendRes)
		if err != nil {
			panic(err)
		}
		if len(friendRes.Friends) == 2 && ultis.CheckEmailValidate(friendRes.Friends[0]) && ultis.CheckEmailValidate(friendRes.Friends[1]) && ultis.CheckFriendConnection(friendRes.Friends[0], friendRes.Friends[1], false) == false {
			db := ultis.DBConn()
			//create new user if not exists
			for _, user := range friendRes.Friends {
				if ultis.CheckUserExisted(user) == false {
					insForm, err := db.Prepare("INSERT INTO USER(email) values(?)")
					if err != nil {
						panic(err)
					}
					insForm.Exec(user)
					fmt.Println("The user '" + user + "' has just been created")
				}
			}
			//insert database
			insForm, err := db.Prepare("INSERT INTO FRIEND_RELATIONSHIP(user1_email,user2_email,blocked) VALUES(?,?,?)")
			if err != nil {
				panic(err)
			}
			insForm.Exec(friendRes.Friends[0], friendRes.Friends[1], 0)
			insForm, err = db.Prepare("INSERT INTO FRIEND_RELATIONSHIP(user1_email,user2_email,blocked) VALUES(?,?,?)")
			if err != nil {
				panic(err)
			}
			insForm.Exec(friendRes.Friends[1], friendRes.Friends[0], 0)

			resultFlag = true
			defer db.Close()
		} else {
			fmt.Println("Error: Invalided data")
		}
		// result := struct {
		// 	Success bool `json:"success"`
		// }{
		// 	resultFlag,
		// }
		var response response.Response
		response.Success = resultFlag
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

// ShowListFriend godoc
// @Summary show list of Friend
// @Description return a list of friend by one user email
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param email path string true "email address"
// @Success 200 {object} models.FriendConnection
// @Router /listFriend [get]
func ShowListFriend(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		emails, ok := r.URL.Query()["email"]

		if !ok || len(emails[0]) < 1 {
			fmt.Println("Url param 'email' is missing")
			return
		}

		email := emails[0]

		var friendResult friendResponse.FriendConnection
		friendResult.Success = false
		if ultis.CheckEmailValidate(email) {
			db := ultis.DBConn()
			friendRecords, err := db.Query("SELECT user2_email from FRIEND_RELATIONSHIP WHERE user1_email='" + email + "' AND blocked=0")
			defer db.Close()
			if err != nil {
				panic(err)
			} else {
				for friendRecords.Next() {
					var email string
					err := friendRecords.Scan(&email)
					if err != nil {
						panic(err)
					}
					friendResult.Friends = append(friendResult.Friends, email)
					friendResult.Count++
				}
				friendResult.Success = true
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Println("res: ", friendResult)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(friendResult); err != nil {
			panic(err)
		}
	}
}

// ShowCommonFriends godoc
// @Summary show common Friends
// @Description return a list of friend by list of user emails
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param emails body models.FriendRequest true "User email list"
// @Success 200 {object} models.FriendConnection
// @Router /commonFriend [post]
func ShowCommonFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var friendRes friendResponse.FriendRequest
		var friendResult friendResponse.FriendConnection
		friendResult.Success = false

		err := json.NewDecoder(r.Body).Decode(&friendRes)
		if err != nil {
			panic(err)
		}
		if len(friendRes.Friends) > 1 {
			db := ultis.DBConn()
			sqlStatement := "select fr1.user2_email from friend_relationship fr1 inner join friend_relationship fr2 on fr1.user1_email='" + friendRes.Friends[0] + "' and fr2.user1_email='" + friendRes.Friends[1] + "' and fr1.user2_email=fr2.user2_email"

			fmt.Println("SQL statement: " + sqlStatement)
			friendRecords, err := db.Query(sqlStatement)
			defer db.Close()
			if err != nil {
				panic(err)
			} else {
				for friendRecords.Next() {
					var email string
					err := friendRecords.Scan(&email)
					if err != nil {
						panic(err)
					}
					friendResult.Friends = append(friendResult.Friends, email)
					friendResult.Count++
				}
				friendResult.Success = true
			}
		} else {
			fmt.Println("Error: Invalided data")
		}
		json.NewEncoder(w).Encode(friendResult)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
