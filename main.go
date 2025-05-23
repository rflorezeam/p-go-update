package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-update/config"
	"github.com/rflorezeam/libro-update/models"
	"github.com/rflorezeam/libro-update/repositories"
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

	var libro models.Libro
	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Formato de libro inválido"})
		return
	}

	result, err := h.service.ActualizarLibro(id, libro)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Libro no encontrado"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Libro actualizado correctamente",
		"result":  result,
	})
}

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros/{id}", handler.ActualizarLibro).Methods("PUT")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	fmt.Printf("Servicio de actualización de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 