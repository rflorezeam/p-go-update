package tests

import (
	"errors"
	"testing"

	"github.com/rflorezeam/libro-update/models"
	"github.com/rflorezeam/libro-update/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroRepository struct {
	mock.Mock
}

func (m *MockLibroRepository) ActualizarLibro(libro *models.Libro) (*models.Libro, error) {
	args := m.Called(libro)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Libro), args.Error(1)
}

func TestActualizarLibro_ServicioExitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	mockRepo.On("ActualizarLibro", libro).Return(libro, nil)

	// Act
	libroActualizado, err := service.ActualizarLibro(libro)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, libroActualizado)
	assert.Equal(t, libro.ID, libroActualizado.ID)
	assert.Equal(t, libro.Titulo, libroActualizado.Titulo)
	assert.Equal(t, libro.Autor, libroActualizado.Autor)
	mockRepo.AssertExpectations(t)
}

func TestActualizarLibro_ServicioError(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("error al actualizar libro en la base de datos")
	mockRepo.On("ActualizarLibro", libro).Return(nil, expectedError)

	// Act
	libroActualizado, err := service.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libroActualizado)
	mockRepo.AssertExpectations(t)
}

func TestActualizarLibro_ServicioLibroNil(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	// Act
	libroActualizado, err := service.ActualizarLibro(nil)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "libro no puede ser nil", err.Error())
	assert.Nil(t, libroActualizado)
	mockRepo.AssertNotCalled(t, "ActualizarLibro")
}

func TestActualizarLibro_ServicioDatosInvalidos(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "", // Título vacío
		Autor:  "Miguel de Cervantes",
	}

	// Act
	libroActualizado, err := service.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "datos de libro inválidos", err.Error())
	assert.Nil(t, libroActualizado)
	mockRepo.AssertNotCalled(t, "ActualizarLibro")
}

func TestActualizarLibro_ServicioIDVacio(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libro := &models.Libro{
		ID:     "", // ID vacío
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	// Act
	libroActualizado, err := service.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "ID no puede estar vacío", err.Error())
	assert.Nil(t, libroActualizado)
	mockRepo.AssertNotCalled(t, "ActualizarLibro")
}

func TestActualizarLibro_ServicioNoEncontrado(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("libro no encontrado")
	mockRepo.On("ActualizarLibro", libro).Return(nil, expectedError)

	// Act
	libroActualizado, err := service.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libroActualizado)
	mockRepo.AssertExpectations(t)
} 