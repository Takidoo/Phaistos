package Includes

import (
	"net/http"
)

func Start() error {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return err
	}
	return nil
}
