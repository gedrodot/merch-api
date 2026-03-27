package handler

import (
	"encoding/json"
	"fmt"
	"merch/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

type authInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth", h.auth)
	mux.HandleFunc("GET /api/users", h.middleware(h.getUsers))

	return mux
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(int)
	users, err := h.service.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-User-Id", fmt.Sprint(userId))
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) auth(w http.ResponseWriter, r *http.Request) {
	var input authInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "неверный формат данных", http.StatusBadRequest)
		return
	}

	jwt, errJwt := h.service.Auth(input.Username, input.Password)
	if errJwt != nil {
		if errJwt.Error() == "service: wrong password" {
			http.Error(w, "неверный пароль", http.StatusUnauthorized)
			return
		}
		http.Error(w, errJwt.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": *jwt,
	})

}
