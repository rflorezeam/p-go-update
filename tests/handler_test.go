package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-update/handlers"
	"github.com/rflorezeam/libro-update/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroService struct {
	mock.Mock
}

func (m *MockLibroService) ActualizarLibro(libro *models.Libro) (*models.Libro, error) {
	args := m.Called(libro)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Libro), args.Error(1)
}

func TestActualizarLibro_Exitoso(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	libroRequest := &models.Libro{
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	libroResponse := &models.Libro{
		ID:     libroID,
		Titulo: libroRequest.Titulo,
		Autor:  libroRequest.Autor,
	}

	mockService.On("ActualizarLibro", mock.MatchedBy(func(l *models.Libro) bool {
		return l.ID == libroID && l.Titulo == libroRequest.Titulo && l.Autor == libroRequest.Autor
	})).Return(libroResponse, nil)

	body, _ := json.Marshal(libroRequest)
	req := httptest.NewRequest(http.MethodPut, "/libros/"+libroID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response models.Libro
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, libroResponse.ID, response.ID)
	assert.Equal(t, libroResponse.Titulo, response.Titulo)
	assert.Equal(t, libroResponse.Autor, response.Autor)
	mockService.AssertExpectations(t)
}

func TestActualizarLibro_NoEncontrado(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	libroRequest := &models.Libro{
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("libro no encontrado")
	mockService.On("ActualizarLibro", mock.MatchedBy(func(l *models.Libro) bool {
		return l.ID == libroID && l.Titulo == libroRequest.Titulo && l.Autor == libroRequest.Autor
	})).Return(nil, expectedError)

	body, _ := json.Marshal(libroRequest)
	req := httptest.NewRequest(http.MethodPut, "/libros/"+libroID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response["error"])
	mockService.AssertExpectations(t)
}

func TestActualizarLibro_Error(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	libroRequest := &models.Libro{
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("error interno del servidor")
	mockService.On("ActualizarLibro", mock.MatchedBy(func(l *models.Libro) bool {
		return l.ID == libroID && l.Titulo == libroRequest.Titulo && l.Autor == libroRequest.Autor
	})).Return(nil, expectedError)

	body, _ := json.Marshal(libroRequest)
	req := httptest.NewRequest(http.MethodPut, "/libros/"+libroID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response["error"])
	mockService.AssertExpectations(t)
}

func TestActualizarLibro_IDVacio(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroRequest := &models.Libro{
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	body, _ := json.Marshal(libroRequest)
	req := httptest.NewRequest(http.MethodPut, "/libros/", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta vacías
	vars := map[string]string{}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ID no proporcionado", response["error"])
}

func TestActualizarLibro_JSONInvalido(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	invalidJSON := []byte(`{"titulo": "El Quijote Actualizado", "autor":}`)
	
	req := httptest.NewRequest(http.MethodPut, "/libros/"+libroID, bytes.NewBuffer(invalidJSON))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "JSON inválido", response["error"])
}

func TestActualizarLibro_DatosInvalidos(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	libroRequest := &models.Libro{
		Titulo: "", // Título vacío
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("datos de libro inválidos")
	mockService.On("ActualizarLibro", mock.MatchedBy(func(l *models.Libro) bool {
		return l.ID == libroID && l.Titulo == libroRequest.Titulo && l.Autor == libroRequest.Autor
	})).Return(nil, expectedError)

	body, _ := json.Marshal(libroRequest)
	req := httptest.NewRequest(http.MethodPut, "/libros/"+libroID, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.ActualizarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response["error"])
	mockService.AssertExpectations(t)
} 