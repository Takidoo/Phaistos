package Database

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) error {
	return nil
}

func DeleteUser(id int) error {
	return nil
}

func DisableUser(id int) error {
	return nil
}

func CheckCredential(username string, password string) (string, error) {
	db_password := ""
	user_id := -1

	err := DB.QueryRow("SELECT id,password FROM users WHERE username=?", username).Scan(&user_id, &db_password)
	if err != nil {
		return "", err
	}

	if !verifyPassword(db_password, password) || user_id == -1 {
		return "", fmt.Errorf("Invalid Password")
	}

	session := generateSessionID(user_id)
	if session == "" {
		return "", err
	}

	return session, nil
}

func CheckSession(session_id string) bool {
	return true
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func verifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateSessionID(id int) string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	session := hex.EncodeToString(b)

	_, err = DB.Query("INSERT INTO sessions (session,user_id) VALUES (?,?)", session, id)

	if err != nil {
		return err.Error()
	}
	return session
}
