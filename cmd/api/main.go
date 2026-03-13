package main

import (
	"github.com/a1bruno/order-api/internal/handlers"
	"github.com/a1bruno/order-api/internal/repositories"
	"github.com/a1bruno/order-api/internal/usecases"
)

func main() {
	repos := repositories.New()
	useCases := usecases.New(repos)
	h := handlers.New(useCases)
	h.Listen(8080)
}
