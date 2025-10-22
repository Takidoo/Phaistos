package API

import (
	"Phaistos/Database"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Username or Password field is empty", http.StatusUnauthorized)
		return
	}

	session_id, err := Database.CheckCredential(username, password)
	if err != nil {
		http.Error(w, "Incorrect Credential", http.StatusUnauthorized)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session_id,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(1 * time.Hour),
	})
}
