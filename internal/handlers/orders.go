package handlers

import (
	"clean-architecture/internal/model"
	"clean-architecture/internal/repository"
	"encoding/json"
	"net/http"
)

type orderHandler struct {
	orderReposito repository.OrderRepository
}

func NewOrderHandler(orderReposito repository.OrderRepository) *orderHandler {
	return &orderHandler{orderReposito}
}

func (o *orderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := o.orderReposito.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (o *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orderDTO model.OrderDTO
	err := json.NewDecoder(r.Body).Decode(&orderDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := o.orderReposito.CreateOrder(orderDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
