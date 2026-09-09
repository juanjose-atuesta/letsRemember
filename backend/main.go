package main

import (
	"backend/handler"
	"backend/services"
	"backend/store"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./lestremember.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	q := `
		CREATE TABLE IF NOT EXISTS INFORMATION (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			theme TEXT NOT NULL,
			description TEXT NOT NULL
		)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	informationStore := store.New(db)
	informationService := services.New(informationStore)
	informationHandler := handler.New(informationService)

	http.HandleFunc("/information", informationHandler.HandlerInformation)
	http.HandleFunc("/information/", informationHandler.HandlerInformationByParamether)

	log.Println("Servidor escuchando en el puerto 4201")
	log.Fatal(http.ListenAndServe(":4201", nil))
}
