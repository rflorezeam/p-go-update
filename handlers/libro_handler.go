package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-update/models"
	"github.com/rflorezeam/libro-update/services"
)

type Handler struct {
	service services.LibroService
}

func NewHandler(service services.LibroService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID no proporcionado"})
		return
	}

	var libro models.Libro
	if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON inválido"})
		return
	}

	libro.ID = id
	libroActualizado, err := h.service.ActualizarLibro(&libro)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "libro no encontrado" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "datos de libro inválidos" {
			statusCode = http.StatusBadRequest
		}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(libroActualizado)
} 