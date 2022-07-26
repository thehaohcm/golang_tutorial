package repositories

import (
	"database/sql"
	"fmt"
	"golang_project/models"
	"golang_project/utils"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type FriendConnectionRepository interface {
	FindFriendsByEmail(email string) []string
	FindCommonFriendsByEmails(emails []string) []string
	CreateFriendConnection(emails []string) bool
	SubscribeFromEmail(req models.SubscribeRequest) bool
	BlockSubscribeByEmail(req models.BlockSubscribeRequest) bool
	GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse
}

type repository struct {
	db *sql.DB
}

func New() FriendConnectionRepository {
	return &repository{
		db: utils.GetInstance(),
	}
}

//1.
func (repo *repository) CreateFriendConnection(emails []string) bool {
	_, err := repo.db.Exec("INSERT INTO friends('user_email','friend_email') VALUES ('" + emails[0] + "','" + emails[1] + "')")

	if err != nil {
		return false
	}
	return true
}

//2.
func (repo *repository) FindFriendsByEmail(email string) []string {
	rows, err := repo.db.Query("SELECT friends FROM friends WHERE user_email=? AND BLOCKED=0", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var friends []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
		}
		friends = append(friends, friend)
	}
	return friends
}

//3.
func (repo *repository) FindCommonFriendsByEmails(emails []string) []string {
	sqlStatement := "SELECT friends FROM friends WHERE"
	for _, email := range emails {
		sqlStatement += " email=" + email + " AND"
	}
	sqlStatement = sqlStatement[:len(sqlStatement)-4]
	sqlStatement += " WHERE BLOCKED=0 GROUP BY friend_email"
	rows, err := repo.db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var friends []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
		}
		friends = append(friends, friend)
	}
	return friends
}

//4.
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) bool {
	_, err := repo.db.Exec("INSERT INTO subscribers('requestor','target') VALUES (?,?)", req.Requestor, req.Target)
	if err != nil {
		panic(err)
		// return false
	}
	return true
}

//5.
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) bool {
	//support A block B:

	if repo.hasFriendConnection(req.Requestor, req.Target) {
		//if A and B are friend, A no longer receive notify from B
		res, err := repo.db.Exec("INSERT OR REPLACE INTO subscribers('requestor','target', 'blocked') VALUES (?,?,1)", req.Requestor, req.Target)
		if err != nil {
			return false
		}

		if _, err = res.LastInsertId(); err != nil {
			return false
		}
	} else {
		//if not friend, no new friend connection added
		res, err := repo.db.Exec("INSERT OR REPLACE INTO friends('user_email','friend_email', 'blocked') VALUES (?,?,1)", req.Requestor, req.Target)
		if err != nil {
			return false
		}
		if _, err = res.LastInsertId(); err != nil {
			panic(err)
		}
	}

	return true
}

func (repo *repository) hasFriendConnection(requestor string, target string) bool {
	rows, err := repo.db.Query("SELECT * FROM friends WHERE user_email=? AND friend_email=? AND BLOCKED=0", requestor, target)
	if err != nil {
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

//6.
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {

	var res models.GetSubscribingEmailListResponse
	res.Success = false

	//if has a friend connection
	rows, err := repo.db.Query("SELECT friend_email FROM friends WHERE user_email=? AND blocked=0", req.Sender)
	if err != nil {
		return res
	}

	defer rows.Close()

	var recipients []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			return res
		}
		recipients = append(recipients, friend)
	}

	//if has a friend connection, but blocked in subscribers tables
	if len(recipients) > 0 {
		rows, err := repo.db.Query("SELECT target FROM subscribers WHERE requestor=? AND blocked=1", req.Sender)
		if err != nil {
			return res
		}

		defer rows.Close()

		var blockedEmails []string
		for rows.Next() {
			var target string
			if err := rows.Scan(&target); err != nil {
				return res
			}
			blockedEmails = append(blockedEmails, target)
		}

		fmt.Println("before: " + strings.Join(recipients, ", "))
		recipients = utils.GetDifference(recipients, blockedEmails)
		fmt.Println("after: " + strings.Join(recipients, ", "))
	}

	//if subscribed to updates
	rows, err = repo.db.Query("SELECT target FROM subscribers WHERE target=? AND blocked=0", req.Sender)
	if err != nil {
		return res
	}

	defer rows.Close()
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			return res
		}
		recipients = append(recipients, friend)
	}

	//if being mentioned in the update
	textArr := strings.Split(req.Text, " ")
	for _, text := range textArr {
		if utils.IsEmailValid(text) {
			recipients = append(recipients, text)
		}
	}

	res.Recipients = recipients
	res.Success = true
	return res
}
