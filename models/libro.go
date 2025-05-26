package models

type Libro struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo string `json:"titulo,omitempty" bson:"titulo,omitempty"`
	Autor  string `json:"autor,omitempty" bson:"autor,omitempty"`
} 