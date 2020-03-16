package ultilities

import (
	"fmt"
	"regexp"
	"strings"
)

func GetEmailList(str string) []string {
	var emails []string
	words := strings.Fields(str)
	for _, word := range words {
		if CheckEmailValidate(word) == true {
			emails = append(emails, word)
		}
	}
	return emails
}

func CheckEmailValidate(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func CheckFriendConnection(user1 string, user2 string, checkBlocked bool) bool {
	result := false
	db := DBConn()
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

func CheckUserExisted(user string) bool {
	result := false
	db := DBConn()
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
