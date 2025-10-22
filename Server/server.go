package Server

import (
	"net/http"
)

func Start() error {
	//Routes
	http.HandleFunc(GAME_ROUTE, GameHandle)
	http.HandleFunc(ADMIN_ROUTE, AdminHandle)
	http.HandleFunc(AUTOR_ROUTE, AutorHandle)
	//Démarrage du serveur
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	return nil
}
