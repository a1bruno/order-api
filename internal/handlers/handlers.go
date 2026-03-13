package handlers

import (
	"fmt"
	"net/http"

	"github.com/a1bruno/order-api/internal/usecases"
)

type Handlers struct {
	useCases usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{}
}

func (h Handlers) Listen(port int) error {
	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}
