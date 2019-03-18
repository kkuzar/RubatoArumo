package users

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/liudng/godump"
	_ "github.com/liudng/godump"
)

var (
	tablename = "users"
	UserKeys  = []string {"id", "username", "family_name",
	"first_name", "password", "mobile", "email", "role",
	"country_id", "city_id", "created_at", "updated_at",
	"last_login", "remark", "isActive", "isDeleted", "isAbnormal",
	"hash_token", "token", "old_token",
	}
	UserAuthDict = []string {"username","password","email", "mobile"}
)

type User struct {
	Uid  	    uint16 			  `db:"id"`
	Username    string 			  `db:"username"`
	Password    string 			  `db:"password"`
	Mobile      string 			  `db:"mobile"`
	Email       string 			  `db:"email"`
	Role        uint8  			  `db:"role"`
	Province_id uint8  			  `db:"province_id"`
	City_id     uint8  			  `db:"city_id"`
	Created_at  mysql.NullTime    `db:"created_at"`
	Updated_at  mysql.NullTime    `db:"updated_at"`
	Last_login  string            `db:"last_login"`
	Last_ip     string            `db:"last_ip"`
	Remark      string 			  `db:"remark"`
	IsActive    bool   			  `db:"isActive"`
	IsDeleted   bool   			  `db:"isDeleted"`
	IsAbnormal  bool   			  `db:"isAbnormal"`
	Hash_token  string 			  `db:"hash_token"`
	Token       string 			  `db:"token"`
	Old_token   string 			  `db:"old_token"`
}

// Data Integrity should be considered in Service Layer (Which means the Validation and Assemble should happend in other layer) .

type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

//func InsertUser(db *sqlx.DB, u User) (sql.Result,error) {
//
//	return  db.Exec(`INSERT INTO users (username, family_name, first_name, password,
//mobile, email, role, country_id, city_id, created_at, updated_at,
//last_login, remark)
//VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
//		u.Username, u.Family_name, u.First_name, u.Password, u.Mobile, u.Email, u.Role, u.Country_id,
//		u.City_id, u.Created_at, u.Updated_at, u.Last_login, u.Remark)
//}

func Test(db *sqlx.DB) {
	var t User
	res  := db.QueryRowx("SELECT * FROM "+tablename+" WHERE id=1")
	res.StructScan(&t)
	godump.Dump(t)
}



//
func InsertUserStruct(db *sqlx.DB, u User ) (sql.Result,error) {
 return db.NamedExec(`INSERT INTO users (username, family_name, first_name, password,
	mobile, email, role, country_id, city_id, created_at, updated_at,
	last_login, remark)
	VALUES (:username, :family_name, :first_name, :password,
	:mobile, :email, :role, :country_id, :city_id, :created_at, :updated_at,
	:last_login, :remark)`,u)
}
//
//func GetUserById(db *sqlx.DB, Userid uint16) (User, error) {
//	var result User
//	res := db.QueryRowx("SELECT * FROM $1 WHERE id=$2",tablename, Userid)
//	res.StructScan(&result)
//	return result, res.Err()
//}
//
//func GetUserByUsername(db *sqlx.DB, username string) (User, error) {
//	var result User
//	res := db.QueryRowx( "SELECT * FROM users WHERE username=$1", username)
//	res.StructScan(&result)
//	return result, res.Err()
//}
//
////func UpdateUserByUsername(db *sqlx.DB, name string) (User, error) {
////	var result User
////}
//
//
//func UpdateUserByUserId()  {
//
//}
//
//func CheckUserExistenceByUsername(db *sqlx.DB, name string) (bool)  {
//	var username string
//	res := db.QueryRow("SELECT username FROM users WHERE username=$1",name)
//	res.Scan(&username)
//	if username == name {
//		return true
//	}
//    return false
//}
//
//func CheckUserExistenceByUserId(db *sqlx.DB, uid uint8) (bool) {
//    var id uint8
//    res := db.QueryRow("SELECT id FROM users WHERE id=$1",uid)
//    res.Scan(&id)
//    if id == uid {
//    	return true
//	}
//	return false
//}
