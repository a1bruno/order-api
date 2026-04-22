package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func NewDBConnection() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco %v", err)
	}

	log.Println("Conexão com banco estabelecida!")
	return db
}
