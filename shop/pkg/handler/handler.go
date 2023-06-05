package handler

import (
	"net/http"

	"github.com/eserzhan/clotheStore/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	// Авторизация
	router.HandleFunc("/auth/sign-up", h.sign_up).Methods("POST")
	router.HandleFunc("/auth/sign-in", h.sign_in)

	// API
	// mux.HandleFunc("/api/lists", h.userIdentity(h.createList)).Methods("POST")
	// mux.HandleFunc("/api/lists", h.userIdentity(h.getList)).Methods("GET")
	// mux.HandleFunc("/api/lists/{id}", h.userIdentity(h.getListById)).Methods("GET")
	// mux.HandleFunc("/api/lists/{id}", h.userIdentity(h.deleteList)).Methods("DELETE")
	// mux.HandleFunc("/api/lists/{id}", h.userIdentity(h.updateListById)).Methods("PUT")

	// mux.HandleFunc("/api/lists/{id}/items", h.userIdentity(h.createItem)).Methods("POST")
	// mux.HandleFunc("/api/lists/{id}/items", h.userIdentity(h.getItems)).Methods("GET")

	// mux.HandleFunc("/api/items/{id}", h.userIdentity(h.getItemById)).Methods("GET")
	// mux.HandleFunc("/api/items/{id}", h.userIdentity(h.deleteItem)).Methods("DELETE")
	// mux.HandleFunc("/api/items/{id}", h.userIdentity(h.updateItem)).Methods("PUT")

	return router
}



func (h *Handler) sign_in(w http.ResponseWriter, r *http.Request) {
	// Обработчик для /auth/sign-in
}

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	// Обработчик для создания списка
}

func (h *Handler) getList(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения списка
}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения списка по ID
}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {
	// Обработчик для удаления списка
}

func (h *Handler) updateListById(w http.ResponseWriter, r *http.Request) {
	// Обработчик для обновления списка по ID
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	// Обработчик для создания элемента
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения элементов
}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	// Обработчик для получения элемента по ID
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	// Обработчик для удаления элемента
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	// Обработчик для обновления элемента
}

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка идентификации пользователя
		next(w, r)
	}
}
