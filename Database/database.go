package Database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Start() {
	var err error
	DB, err = sql.Open("sqlite3", "Database/phaistos.db")
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données:", err)
	}

	fmt.Println("Connexion à la base de données réussie")
	CreateTables()
}

func Stop() {
	if DB != nil {
		DB.Close()
		fmt.Println("Connexion à la base de données fermée")
	}
}

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			register TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			account_disabled BOOLEAN DEFAULT false
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Erreur lors de la création des tables de la base de données:", err)
		}
	}
	fmt.Println("Base de donnée démarrée avec succès")
}
