package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/google/uuid"
	"log"
	"net/http"
)

type User struct {
	Id            string `json:"id"`
	User_name     string `json:"user_name"`
	User_email    string `json:"user_email"`
	User_password string `json:"user_password"`
	Created_at    string
}
type idJson struct {
	Id string
}

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS "User"(
		id VARCHAR(129) PRIMARY KEY,
    	user_name VARCHAR(25) NOT NULL,
		user_email VARCHAR(40) UNIQUE NOT NULL,
		password VARCHAR(25) NOT NULL,
		created_at timestamp DEFAULT NOW()
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("couldn't create users table:%v", err)
	}
}
func InsertInUserTable(w http.ResponseWriter, db *sql.DB, user *User) {
	if user.User_name == "" {
		AnswerWithError(w, 400, fmt.Sprintf("user_name field is empty"))
		return
	}
	if user.User_password == "" {
		AnswerWithError(w, 400, fmt.Sprintf("password field is empty"))
		return
	}
	if user.User_email == "" {
		AnswerWithError(w, 400, fmt.Sprintf("user_email field is empty"))
		return
	}
	query := `INSERT INTO "User"(id,user_name,user_email,password)
	VALUES ($1,$2,$3,$4) RETURNING id`
	var id string
	password := base64.StdEncoding.EncodeToString([]byte(user.User_password))
	email := base64.StdEncoding.EncodeToString([]byte(user.User_email))
	encodedId := password + "/" + email
	err := db.QueryRow(query, encodedId, user.User_name, user.User_email, user.User_password).Scan(&id)
	if err != nil {
		AnswerWithError(w, 400, fmt.Sprintf("couldn't insert User data in database: %v", err))
		return
	}
	AnswerWithJson(w, 200, idJson{Id: id})
}

//что бы задекодить делай это
//decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
//decodedString := string(decodedBytes)
//не забудь разделить

// -----------------------------------------------------------------
func GetUserData(w http.ResponseWriter, db *sql.DB, user *User) {
	query := `SELECT user_name,user_email,password,to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') FROM "User" 
    WHERE id=$1`
	var (
		user_name  string
		password   string
		user_email string
		created_at string
	)
	err := db.QueryRow(query, user.Id).Scan(&user_name, &password, &user_email, &created_at)
	if err != nil {
		AnswerWithError(w, 400, fmt.Sprintf("wrong id key or ivalid format: %v", err))
		return
	}
	AnswerWithJson(w, 200, &User{
		Id:            user.Id,
		User_name:     user_name,
		User_email:    password,
		User_password: user_email,
		Created_at:    created_at,
	})
}
