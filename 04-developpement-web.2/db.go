package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func ConnectDB(dataSource string) error {
	log.Println("Connexion à la base de données")
	if db != nil {
		// Already connected
		log.Println("Déjà connecté à la base de données")
		return nil
	}
	var err error
	db, err = sql.Open("postgres", dataSource)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("Connecté à la base de données")
	return nil
}

func InitDB(file string) error {
	log.Println("Initialisation de la base de données")
	err := ConnectDB(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName))
	if err != nil {
		return err
	}
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(fileContent))
	if err != nil {
		return err
	}
	log.Println("Base de données initialisée")
	return db.Close()
}
