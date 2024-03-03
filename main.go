package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"streamer/webapp/api/auth"
	"streamer/webapp/db"
)

func main() {
	r := mux.NewRouter()

	db.SetupConnection()
	defer db.CloseConnection()

	if err := db.MigrateModels(db.Conn); err != nil {
		log.Fatal(err)
	}

	auth.SetupRouter(r)

	log.Println("Started Server")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal("Can not start server")
	}
}
