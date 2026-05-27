package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name string `bson:"name"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	Avatar string `bson:"avatar"`
	Status string `bson:"status"`
	CreatedAt time.Time `bson:"createdAt"`
}