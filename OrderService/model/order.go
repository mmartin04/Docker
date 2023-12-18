package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name  string             `json:"name"`
    // Weitere Felder nach Bedarf...
}
