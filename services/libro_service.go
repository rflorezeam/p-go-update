package services

import (
	"github.com/rflorezeam/libro-update/models"
	"github.com/rflorezeam/libro-update/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroService interface {
	ActualizarLibro(id string, libro models.Libro) (*mongo.UpdateResult, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) ActualizarLibro(id string, libro models.Libro) (*mongo.UpdateResult, error) {
	return s.repo.ActualizarLibro(id, libro)
} 