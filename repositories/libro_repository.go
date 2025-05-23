package repositories

import (
	"context"

	"github.com/rflorezeam/libro-update/config"
	"github.com/rflorezeam/libro-update/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	ActualizarLibro(id string, libro models.Libro) (*mongo.UpdateResult, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{
		collection: config.GetCollection(),
	}
}

func (r *libroRepository) ActualizarLibro(id string, libro models.Libro) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"titulo": libro.Titulo,
			"autor":  libro.Autor,
		},
	}

	return r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
} 