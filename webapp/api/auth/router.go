package auth

import "github.com/gorilla/mux"

func SetupRouter(r *mux.Router) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", login).Methods("POST")
	authRouter.HandleFunc("/register", register).Methods("POST")
}
