package Server

import (
	"net/http"
)

func Start() error {
	//Route vers le jeu
	http.HandleFunc(GAME_ROUTE, nil)

	//Démarrage du serveur
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	return nil
}
