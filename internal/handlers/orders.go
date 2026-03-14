package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/a1bruno/order-api/internal/models"
	"github.com/google/uuid"
)

func (h Handlers) registerOrderEndpoints() {
	http.HandleFunc("GET /orders", h.getAllOrders)
	http.HandleFunc("POST /orders", h.addOrder)
	http.HandleFunc("PATCH /orders/{id}/refund", h.refund)
}

func (h Handlers) getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.useCases.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h Handlers) addOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.CreateOrder //cria a variavel com a struct que recebe apenas os parametros necessarios

	json.NewDecoder(r.Body).Decode(&newOrder) //cria um decoder para ler a request do usuario//decodifica o json do body para atribuir as variaveis no var

	id := h.useCases.Add(newOrder) //passa a variavel para o usecaso que faz a implementação necessária

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ResponseOrder{NewOrderID: id}) //monta a struct de resposta ao usuario com o id gerado e serializa em json
}

func (h Handlers) refund(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")   // atribui o valor do id na requisiçào do usuario
	id, err := uuid.Parse(idStr) //transforma o valor capturado em uuid e retorna um erro, se o id for valido o erro é nulo, caso contrario é != nulo.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //se o erro for diferente de nil, vamos retornar o bad request e breakar a continuidade
		return
	}
	h.useCases.Refund(id)
	w.WriteHeader(http.StatusNoContent)
}
