package main

import (
	"Phaistos/Server"
)

func StartApp() error {
	err := Server.Start()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := StartApp()
	if err != nil {

	}
	print("Impossible de démarrer l'application web, erreur : ", err.Error())
}
