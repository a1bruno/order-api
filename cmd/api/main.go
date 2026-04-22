package main

import (
	"log"

	"github.com/a1bruno/order-api/db"
	"github.com/a1bruno/order-api/internal/handlers"
	"github.com/a1bruno/order-api/internal/repositories"
	"github.com/a1bruno/order-api/internal/usecases"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env")
	}

	conn := db.NewDBConnection()
	defer conn.Close()
	repos := repositories.New()
	useCases := usecases.New(repos)
	h := handlers.New(useCases)
	h.Listen(8080)
}
