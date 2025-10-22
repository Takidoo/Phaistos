package Server

import (
	"net/http"
)

func Start() error {
	//Routes
	http.HandleFunc(GAME_ROUTE, GameHandle)

	//Démarrage du serveur
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	return nil
}
