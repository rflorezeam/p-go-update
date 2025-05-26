package services

import (
	"errors"

	"github.com/rflorezeam/libro-update/models"
	"github.com/rflorezeam/libro-update/repositories"
)

type LibroService interface {
	ActualizarLibro(libro *models.Libro) (*models.Libro, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) ActualizarLibro(libro *models.Libro) (*models.Libro, error) {
	if libro == nil {
		return nil, errors.New("libro no puede ser nil")
	}

	if libro.ID == "" {
		return nil, errors.New("ID no puede estar vacío")
	}

	if libro.Titulo == "" || libro.Autor == "" {
		return nil, errors.New("datos de libro inválidos")
	}

	return s.repo.ActualizarLibro(libro)
} 