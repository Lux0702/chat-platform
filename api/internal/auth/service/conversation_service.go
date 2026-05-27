package service

import (
	"chat-platform-api/internal/auth/entity"
	"chat-platform-api/internal/auth/repository"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConversationService struct {
	repo *repository.ConversationRepository
}

func NewConversationService() *ConversationService {

	return &ConversationService{
		repo: repository.NewConversationRepository(),
	}
}

func (s *ConversationService) CreatePrivateConversation(
	currentUserID string,
	targetUserID string,
) error {
	currentObjectID, _ := primitive.ObjectIDFromHex(currentUserID)
	targetObjectID, _ := primitive.ObjectIDFromHex(targetUserID)

	existing, _ := s.repo.FindPrivateConversation(
		currentObjectID,
		targetObjectID,
	)

	if existing != nil {
		return errors.New(
			"conversation already exists",
		)
	}
	conversation := &entity.Conversation{
		Type: "private",
		Members: []primitive.ObjectID{
			currentObjectID,
			targetObjectID,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Create(conversation)
}
