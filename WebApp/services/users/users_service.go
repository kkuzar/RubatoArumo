package users_service

import (
	"github.com/huzhaer/qianxun/model/users"
	"github.com/jmoiron/sqlx"
	"html"
	"strings"
)

func checkOrInit (key string ,u map[string]string) (string) {
	value, ok := u[key]
	if ok == false {
		u[key] = ""
	}
	return removeTag(value)
}

func removeTag(str string) string{
	return html.EscapeString(strings.Replace(str,"</","",-1))
}

func GetPassWordWithUserName(username string, db *sqlx.DB ) string  {
	var password string
	res := db.QueryRowx("SELECT password FROM users WHERE username=?", username)
	res.Scan(&password)
	return password
}

func GetUserWithUserName(username string ,  db *sqlx.DB) users.User  {
	var t users.User
	res := db.QueryRowx("SELECT * FROM users WHERE username=?", username)
	res.Scan(&t)
	return t
}

// func InsertNewUserOrUpdateExistOnes()