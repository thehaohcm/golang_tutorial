package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	commonsResponse "Assigment/models/commons"
	subscribeResponse "Assigment/models/subscribe"
	ultis "Assigment/ultilities"
)

// Subscribe godoc
// @Summary subscribe user
// @Description return a result of subscribing user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param subscribe body models.SubscribeRequest true "Subscribe"
// @Success 200 {object} models.Response
// @Router /subscribe [post]
func Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var subscribeRes subscribeResponse.SubscribeRequest
		var response commonsResponse.Response
		response.Success = false
		err := json.NewDecoder(r.Body).Decode(&subscribeRes)
		if err != nil {
			panic(err)
		} else {
			if ultis.CheckEmailValidate(subscribeRes.Requestor) && ultis.CheckEmailValidate(subscribeRes.Target) {
				db := ultis.DBConn()
				insForm, err := db.Prepare("INSERT INTO SUBSCRIBE(requester_email,target_email,blocked) VALUES(?,?,?)")
				if err != nil {
					panic(err)
				}
				insForm.Exec(subscribeRes.Requestor, subscribeRes.Target, 0)
				response.Success = true
				defer db.Close()
			}
		}

		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// BlockSubscribe godoc
// @Summary block Subscribe user
// @Description return a result of blocking subscribe user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param subscribe body models.SubscribeRequest true "Subscribe"
// @Success 200 {object} models.Response
// @Router /blockSubscribe [post]
func BlockSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var subscribeRes subscribeResponse.SubscribeRequest
		var response commonsResponse.Response
		response.Success = false
		err := json.NewDecoder(r.Body).Decode(&subscribeRes)
		if err != nil {
			panic(err)
		} else {
			db := ultis.DBConn()

			//check if they are friends or not
			//not yet been implemented
			var updateForm *sql.Stmt
			var err error
			if ultis.CheckFriendConnection(subscribeRes.Requestor, subscribeRes.Target, false) {
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

// ShowListRecipients godoc
// @Summary Show List of Recipients
// @Description get list by email user
// @Tags Subscribe
// @Accept  json
// @Produce  json
// @Param recipients body models.RecipientRequest true "Recipient"
// @Success 200 {object} models.RecipientResponse
// @Router /listRecipients [post]
func ShowListRecipients(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var recipientReq subscribeResponse.RecipientRequest
		var recipientRes subscribeResponse.RecipientResponse
		recipientRes.Success = false
		err := json.NewDecoder(r.Body).Decode(&recipientReq)
		if err != nil {
			panic(err)
		} else {
			emailList := ultis.GetEmailList(recipientReq.Text)
			if emailList != nil && len(emailList) > 0 {
				db := ultis.DBConn()
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
