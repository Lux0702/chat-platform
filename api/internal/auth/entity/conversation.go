package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LastMessage struct{
	Content string `bson:"content"`
	SenderID string `bson:"senderId"`
	CreatedAt time.Time `bson:"createdAt"`
}
type Conversation struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Type string `bson:"type"`
	Name string `bson:"name,omitempty"`
	Avatar string `bson:"avatar,omitempty"`
	Members []primitive.ObjectID `bson:"members"`
	LastMessage LastMessage `bson:"lastMessage,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}