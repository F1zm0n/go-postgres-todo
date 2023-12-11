package Db

import (
	"database/sql"
	"encoding/base64"
	_ "github.com/google/uuid"
	"log"
)

type User struct {
	User_name     string `json:"user_name"`
	User_email    string `json:"user_email"`
	User_password string `json:"user_password"`
}

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "User"(
		id VARCHAR(129) PRIMARY KEY,
    	user_name VARCHAR(25) NOT NULL,
		User_email VARCHAR(40) UNIQUE NOT NULL,
		password VARCHAR(25) NOT NULL,
		created_at timestamp DEFAULT NOW()
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("couldn't create users table:%v", err)
	}
}
func InsertInUserTable(db *sql.DB, user User) string {
	query := `INSERT INTO User(id,user_name,User_email,password)
	VALUES ($1,$2,$3,$4) RETURNING id`
	var id string
	password := base64.StdEncoding.EncodeToString([]byte(user.User_password))
	email := base64.StdEncoding.EncodeToString([]byte(user.User_email))
	encodedId := password + "/" + email
	err := db.QueryRow(query, encodedId, user.User_name, user.User_email, user.User_password).Scan(&id)
	if err != nil {
		log.Println("couldn't insert User data in database: ", err)
		return ""
	}
	return id
}

//что бы задекодить делай это
//decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
//decodedString := string(decodedBytes)
//не забудь разделить

// -----------------------------------------------------------------
func GetUserData(db *sql.DB) User {
	query := `SELECT user_name,password,User_email FROM User
	VALUES($1) WHERE id=$1 RETURNING user_name,password,User_email `
	var (
		user_name  string
		password   string
		user_email string
	)
	err := db.QueryRow(query).Scan(&user_name, &password, &user_email)
	if err != nil {
		log.Println("wrong id key or ivalid format")
		return User{}
	}
	return User{
		User_name:     user_name,
		User_email:    password,
		User_password: user_email,
	}

}