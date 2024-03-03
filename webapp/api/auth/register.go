package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"streamer/webapp/auth"
	"streamer/webapp/db"
	"streamer/webapp/db/models"
)

type registerBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rb *registerBody) isPasswordCorrect() bool {
	if rb.Password == "" {
		return false
	}
	return true
}

func register(w http.ResponseWriter, r *http.Request) {
	var body registerBody

	errBody := json.NewDecoder(r.Body).Decode(&body)
	if errBody != nil {
		http.Error(w, errBody.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(body)
	if !body.isPasswordCorrect() {
		http.Error(w, "Incorrect password", http.StatusFailedDependency)
		return
	}

	hashedPassword := auth.HashPassword(body.Password)

	user := &models.User{
		Email:    body.Email,
		Password: hashedPassword,
	}

	_, errInsert := db.Conn.Model(user).Insert()
	if errInsert != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Created user with Id: %d, Email: %s", user.Id, user.Email)
}
