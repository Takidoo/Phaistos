package main

import (
	"Phaistos/Database"
	"Phaistos/Server"
)

func StartApp() error {
	Database.Start()
	err := Server.Start()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := StartApp()
	if err != nil {
		print("Impossible de démarrer l'application web, erreur : ", err.Error())
	}
}
