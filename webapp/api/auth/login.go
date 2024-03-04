package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"streamer/webapp/auth"
	"streamer/webapp/db"
	"streamer/webapp/db/models"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var body loginBody

	errBody := json.NewDecoder(r.Body).Decode(&body)
	if errBody != nil {
		log.Printf("%s", errBody.Error())
		http.Error(w, errBody.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword := auth.HashPassword(body.Password)

	user := new(models.User)
	selectErr := db.Conn.Model(user).
		Where("\"user\".email = ? AND \"user\".password = ?", body.Email, hashedPassword).
		Select()

	if selectErr != nil {
		log.Printf("%s", selectErr.Error())
		http.Error(w, selectErr.Error(), http.StatusUnauthorized)
		return
	}

	token, errGenerateToken := auth.GenerateToken(user)
	if errGenerateToken != nil {
		log.Printf("%s", errGenerateToken.Error())
		http.Error(w, errGenerateToken.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(token)
	return
}
