package handler

import "net/http"

func (h *Handler) sign_up(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Запись текста "Hello, World!" в тело ответа
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		// Обработка ошибки записи
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}