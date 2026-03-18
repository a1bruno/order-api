package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/a1bruno/order-api/internal/models"
	"github.com/a1bruno/order-api/internal/repositories/orders"
	"github.com/google/uuid"
)

func (h Handlers) registerOrderEndpoints() {
	http.HandleFunc("GET /orders", h.getAllOrders)
	http.HandleFunc("POST /orders", h.addOrder)
	http.HandleFunc("PATCH /orders/{id}/refund", h.refund)
}

func (h Handlers) getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.useCases.GetAll()
	if err != nil {
		log.Printf("GetAll falhou: %v", err)
		http.Error(w, "erro interno ao consultar pedidos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h Handlers) addOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.CreateOrder

	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Formato inválido.", http.StatusBadRequest)
		return
	}
	id, err := h.useCases.Add(newOrder)
	if err != nil {
		log.Printf("addOrder falhou %v", err)
		http.Error(w, "Erro ao criar pedido", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ResponseOrder{NewOrderID: id})
}

func (h Handlers) refund(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.useCases.Refund(id)
	if err != nil {
		if errors.Is(err, orders.ErrNotFound) {
			http.Error(w, "Pedido não encontrado", http.StatusNotFound)
			return
		}
		if errors.Is(err, orders.ErrOrderAlreadyRefunded) {
			http.Error(w, "pedido já estornado", http.StatusConflict)
			return
		}
		log.Printf("refund falhou %v", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
