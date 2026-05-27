package repository

import (
	"context"
	"chat-platform-api/internal/auth/entity"
	"chat-platform-api/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
	 "go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{
	Collection *mongo.Collection
}
func NewUserRepository() *UserRepository{
	return &UserRepository{
		Collection: database.DB.Collection("users"),
	}
}

func (r *UserRepository) FindByEmail(email string,) (*entity.User, error){
		var user entity.User

		err:= r.Collection.FindOne(
			context.Background(),
			bson.M{
				"email": email,
			},
		).Decode(&user)

		return &user, err
}

func (r *UserRepository) Create(user *entity.User) error{
	_, err:= r.Collection.InsertOne(
		context.Background(),
		user,
	)
	return err
}
