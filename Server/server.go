package Server

import (
	"Phaistos/Database"
	"Phaistos/Server/API"
	"database/sql"
	"net/http"
)

func Start() error {
	//Routes
	http.HandleFunc(GAME_ROUTE, GameHandle)
	http.HandleFunc(ADMIN_ROUTE, AdminHandle)
	http.HandleFunc(USER_ROUTE, ProfileHandle)
	http.HandleFunc("/test", Testhandle)

	//APIs
	http.HandleFunc("/api/login", API.Login)

	err := createDefaultAdmin(Database.DB)
	if err != nil {
		return err
	}

	//Démarrage du serveur
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	return nil
}

func createDefaultAdmin(db *sql.DB) error {
	print("No error")
	username := ADMIN_USERNAME
	password := ADMIN_PASSWORD

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	hashedPwd := Database.HashPassword(password)

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPwd)
	if err != nil {
		return err
	}
	print("No error")
	return nil
}
