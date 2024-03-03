package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var Conn *pg.DB

func SetupConnection() {
	Conn = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "main_db",
		Addr:     "localhost:5432",
	})
}

func CloseConnection() {
	err := Conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
