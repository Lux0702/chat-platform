package repository

import (
	"chat-platform-api/internal/auth/entity"
	"chat-platform-api/pkg/database"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	 "go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/bson"
)

type ConversationRepository struct {
	Collection *mongo.Collection
}

func NewConversationRepository() *ConversationRepository{
	return &ConversationRepository{
		Collection: database.DB.Collection("conversations"),
	}
}

func (r *ConversationRepository) Create(conversation *entity.Conversation) error{

	_, err := r.Collection.InsertOne(
		context.Background(),
		conversation,
	)
	return err
	
}
func (r *ConversationRepository) FindPrivateConversation(
	userA primitive.ObjectID,
	userB primitive.ObjectID,
) (*entity.Conversation, error){
 
	var conversation entity.Conversation

	filter:= bson.M{
		"type": "private",
		"members": bson.M{
			"$all": []primitive.ObjectID{userA, userB},
		},
	}

	err := r.Collection.FindOne(
		context.Background(),
		filter,
	).Decode(&conversation)
	if err !=nil{
		return nil, err
	}
		return &conversation, nil
}