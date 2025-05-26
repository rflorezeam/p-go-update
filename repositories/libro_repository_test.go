package repositories

import (
	"errors"
	"testing"

	"github.com/rflorezeam/libro-update/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
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

func TestActualizarLibro_RepositorioExitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	mockRepo.On("ActualizarLibro", libro).Return(libro, nil)

	// Act
	libroActualizado, err := mockRepo.ActualizarLibro(libro)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, libroActualizado)
	assert.Equal(t, libro.ID, libroActualizado.ID)
	assert.Equal(t, libro.Titulo, libroActualizado.Titulo)
	assert.Equal(t, libro.Autor, libroActualizado.Autor)
	mockRepo.AssertExpectations(t)
}

func TestActualizarLibro_RepositorioError(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("error de conexión con la base de datos")
	mockRepo.On("ActualizarLibro", libro).Return(nil, expectedError)

	// Act
	libroActualizado, err := mockRepo.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libroActualizado)
	mockRepo.AssertExpectations(t)
}

func TestActualizarLibro_RepositorioIDInvalido(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libro := &models.Libro{
		ID:     "invalid-id",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := errors.New("ID inválido")
	mockRepo.On("ActualizarLibro", libro).Return(nil, expectedError)

	// Act
	libroActualizado, err := mockRepo.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libroActualizado)
	mockRepo.AssertExpectations(t)
}

func TestActualizarLibro_RepositorioNoEncontrado(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libro := &models.Libro{
		ID:     "123",
		Titulo: "El Quijote Actualizado",
		Autor:  "Miguel de Cervantes",
	}

	expectedError := mongo.ErrNoDocuments
	mockRepo.On("ActualizarLibro", libro).Return(nil, expectedError)

	// Act
	libroActualizado, err := mockRepo.ActualizarLibro(libro)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, libroActualizado)
	mockRepo.AssertExpectations(t)
} 