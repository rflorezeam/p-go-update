package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-update/config"
	"github.com/rflorezeam/libro-update/handlers"
	"github.com/rflorezeam/libro-update/repositories"
	"github.com/rflorezeam/libro-update/services"
)

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := handlers.NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros/{id}", handler.ActualizarLibro).Methods("PUT")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	fmt.Printf("Servicio de actualización de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 