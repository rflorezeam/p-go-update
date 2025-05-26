package repositories

import (
	"context"
	"errors"

	"github.com/rflorezeam/libro-update/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LibroRepository interface {
	ActualizarLibro(libro *models.Libro) (*models.Libro, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{}
}

func (r *libroRepository) ActualizarLibro(libro *models.Libro) (*models.Libro, error) {
	if libro == nil {
		return nil, errors.New("libro no puede ser nil")
	}

	objectID, err := primitive.ObjectIDFromHex(libro.ID)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"titulo": libro.Titulo,
			"autor":  libro.Autor,
		},
	}

	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := r.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		&opts,
	)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("libro no encontrado")
		}
		return nil, result.Err()
	}

	var libroActualizado models.Libro
	if err := result.Decode(&libroActualizado); err != nil {
		return nil, err
	}

	return &libroActualizado, nil
} 