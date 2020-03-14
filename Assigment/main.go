package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"

	_ "./docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type user struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type friendRequest struct {
	Friends []string `json:"friends"`
}

type friendConnection struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type subscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type recipientRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type recipientResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

type response struct {
	Success bool `json:"success"`
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

// showAllUser godoc
// @Summary show all infos of all users
// @Description return a list of entire users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} user
// @Router /showAllUsers [get]
func showAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := dbConn()
		selDb, err := db.Query("SELECT email, first_name, last_name from user")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var res []user
		for selDb.Next() {
			var user user
			var firstName, lastName sql.NullString
			err := selDb.Scan(&user.Email, &firstName, &lastName)
			if firstName.Valid {
				user.FirstName = firstName.String
			}
			if lastName.Valid {
				user.LastName = lastName.String
			}
			if err != nil {
				panic(err.Error())
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

// // show godoc
// // @Summary add a new friend
// // @Description return a result of creating a new friend by 2 user emails
// // @Tags Friend
// // @Accept  json
// // @Produce  json
// // @Param friend body friendRequest true "friend"
// // @Success 200 {object} response
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

// addFriend godoc
// @Summary add a new friend
// @Description return a result of creating a new friend by 2 user emails
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param friend body friendRequest true "friend"
// @Success 200 {object} response
// @Router /addFriend [post]
func addFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		resultFlag := false

		var friendRes friendRequest

		err := json.NewDecoder(r.Body).Decode(&friendRes)
		if err != nil {
			panic(err)
		}
		if len(friendRes.Friends) == 2 && checkEmailValidate(friendRes.Friends[0]) && checkEmailValidate(friendRes.Friends[1]) && checkFriendConnection(friendRes.Friends[0], friendRes.Friends[1], false) == false {
			db := dbConn()
			//create new user if not exists
			for _, user := range friendRes.Friends {
				if checkUserExisted(user) == false {
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
		var response response
		response.Success = resultFlag
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

// showListFriend godoc
// @Summary show list of Friend
// @Description return a list of friend by one user email
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param email query string true "email address"
// @Success 200 {object} friendConnection
// @Router /listFriend [get]
func showListFriend(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		emails, ok := r.URL.Query()["email"]

		if !ok || len(emails[0]) < 1 {
			fmt.Println("Url param 'email' is missing")
			return
		}

		email := emails[0]

		var friendResult friendConnection
		friendResult.Success = false

		db := dbConn()
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

		w.WriteHeader(http.StatusOK)
		fmt.Println("res: ", friendResult)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(friendResult); err != nil {
			panic(err)
		}
	}
}

// showCommonFriends godoc
// @Summary show common Friends
// @Description return a list of friend by list of user emails
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param emails body friendRequest true "User email list"
// @Success 200 {object} friendConnection
// @Router /commonFriend [post]
func showCommonFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var friendRes friendRequest
		var friendResult friendConnection
		friendResult.Success = false

		err := json.NewDecoder(r.Body).Decode(&friendRes)
		if err != nil {
			panic(err)
		}
		if len(friendRes.Friends) > 1 {
			db := dbConn()
			sqlStatement := "SELECT user2_email from FRIEND_RELATIONSHIP WHERE "
			var firstClause, secondClause string
			for i := 0; i < len(friendRes.Friends); i++ {
				//sqlStatement = sqlStatement + "user1_email='" + friendRes.Friends[i] + "' AND user2_email <> '" + friendRes.Friends[i] + "'"
				firstClause = firstClause + "user1_email='" + friendRes.Friends[i] + "'"
				secondClause = secondClause + "user2_email<>'" + friendRes.Friends[i] + "'"
				if i < len(friendRes.Friends)-1 {
					firstClause = firstClause + " OR "
					secondClause = secondClause + " AND "
				}
			}
			sqlStatement = sqlStatement + "(" + firstClause + ") AND (" + secondClause + ")"
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

// subscribe godoc
// @Summary subscribe user
// @Description return a result of subscribing user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param subscribe body subscribeRequest true "Subscribe"
// @Success 200 {object} response
// @Router /subscribe [post]
func subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var subscribeRes subscribeRequest
		var response response
		response.Success = false
		err := json.NewDecoder(r.Body).Decode(&subscribeRes)
		if err != nil {
			panic(err)
		} else {
			db := dbConn()
			insForm, err := db.Prepare("INSERT INTO SUBSCRIBE(requester_email,target_email,blocked) VALUES(?,?,?)")
			if err != nil {
				panic(err)
			}
			insForm.Exec(subscribeRes.Requestor, subscribeRes.Target, 0)
			response.Success = true
			defer db.Close()
		}

		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// blockSubscribe godoc
// @Summary block Subscribe user
// @Description return a result of blocking subscribe user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param subscribe body subscribeRequest true "Subscribe"
// @Success 200 {object} response
// @Router /blockSubscribe [post]
func blockSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var subscribeRes subscribeRequest
		var response response
		response.Success = false
		err := json.NewDecoder(r.Body).Decode(&subscribeRes)
		if err != nil {
			panic(err)
		} else {
			db := dbConn()

			//check if they are friends or not
			//not yet been implemented
			var updateForm *sql.Stmt
			var err error
			if checkFriendConnection(subscribeRes.Requestor, subscribeRes.Target, false) {
				updateForm, err := db.Prepare("UPDATE FRIEND_RELATIONSHIP SET BLOCKED = 1 WHERE (user1_email= ? AND user2_email = ?) OR (user1_email= ? AND user2_email=?)")
				if err != nil {
					panic(err)
				}
				updateForm.Exec(subscribeRes.Requestor, subscribeRes.Target, subscribeRes.Target, subscribeRes.Requestor)
			}

			updateForm, err = db.Prepare("UPDATE SUBSCRIBE SET BLOCKED = 1 WHERE requester_email= ? AND target_email=?")
			if err != nil {
				panic(err)
			}
			updateForm.Exec(subscribeRes.Requestor, subscribeRes.Target)
			response.Success = true
			defer db.Close()
		}

		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// showListRecipients godoc
// @Summary Show List of Recipients
// @Description get list by email user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param recipients body recipientRequest true "Recipient"
// @Success 200 {object} recipientResponse
// @Router /listRecipients [post]
func showListRecipients(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var recipientReq recipientRequest
		var recipientRes recipientResponse
		recipientRes.Success = false
		err := json.NewDecoder(r.Body).Decode(&recipientReq)
		if err != nil {
			panic(err)
		} else {
			emailList := getEmailList(recipientReq.Text)
			if emailList != nil && len(emailList) > 0 {
				db := dbConn()
				sqlStatment := "SELECT requester_email FROM SUBSCRIBE WHERE target_email ='" + recipientReq.Sender + "'" //+ "' and blocked=0"
				resultRecords, err := db.Query(sqlStatment)
				if err != nil {
					panic(err)
				}
				for resultRecords.Next() {
					var email string
					err := resultRecords.Scan(&email)
					if err != nil {
						panic(err)
					}
					emailList = append(emailList, email)
				}
				sqlStatment = "SELECT user2_email FROM FRIEND_RELATIONSHIP WHERE user1_email = '" + recipientReq.Sender + "' and blocked=0"
				resultRecords, err = db.Query(sqlStatment)
				if err != nil {
					panic(err)
				}
				for resultRecords.Next() {
					var email string
					err := resultRecords.Scan(&email)
					if err != nil {
						panic(err)
					}
					emailList = append(emailList, email)
				}
				if len(emailList) > 0 {
					recipientRes.Success = true
					recipientRes.Recipients = emailList
				}
			}
		}
		json.NewEncoder(w).Encode(recipientRes)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getEmailList(str string) []string {
	var emails []string
	words := strings.Fields(str)
	for _, word := range words {
		if checkEmailValidate(word) == true {
			emails = append(emails, word)
		}
	}
	return emails
}

func checkEmailValidate(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func checkFriendConnection(user1 string, user2 string, checkBlocked bool) bool {
	result := false
	db := dbConn()
	statement := "SELECT exists(SELECT user1_email FROM Friend_Relationship WHERE user1_email = '" + user1 + "' and user2_email = '" + user2 + "'"
	if checkBlocked == true {
		statement += " and blocked = 0"
	}
	statement += ")"
	fmt.Println("SQL statement: " + statement)
	var exists bool
	row := db.QueryRow(statement)
	if err := row.Scan(&exists); err != nil {
		result = true
	}
	defer db.Close()
	return result
}

func checkUserExisted(user string) bool {
	result := false
	db := dbConn()
	statement := "SELECT exists(SELECT email FROM USER WHERE email = '" + user + "')"
	fmt.Println("SQL statement: " + statement)
	var exists bool
	row := db.QueryRow(statement)
	if err := row.Scan(&exists); err == nil {
		result = exists
	} else {
		panic(err)
	}
	defer db.Close()
	return result
}

// @title Golang API endpoints assignment
// @version 1.0
// @description This is a sample of Golang API endpoints assignment
// @termsOfService http://swagger.io/terms/

// @contact.name Hao Nguyen
// @contact.url http://musicmaven.s3corp.vn
// @contact.email hao.nguyen@s3corp.com.vn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	log.Println("Server started on: http://localhost:8080")

	// http.HandleFunc("/showAllUsers", showAllUser)
	// // http.HandleFunc("/show", show)
	// http.HandleFunc("/addFriend", addFriend)
	// http.HandleFunc("/listFriend", showListFriend)
	// http.HandleFunc("/commonFriend", showCommonFriends)
	// http.HandleFunc("/subscribe", subscribe)
	// http.HandleFunc("/blockSubscribe", blockSubscribe)
	// http.HandleFunc("/listRecipients", showListRecipients)

	//swagger
	router := chi.NewRouter()

	router.Get("/showAllUsers", showAllUser)
	router.Post("/addFriend", addFriend)
	router.Get("/listFriend", showListFriend)
	router.Post("/commonFriend", showCommonFriends)
	router.Post("/subscribe", subscribe)
	router.Post("/blockSubscribe", blockSubscribe)
	router.Post("/listRecipients", showListRecipients)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition"
	))

	http.ListenAndServe(":8080", router)
}
