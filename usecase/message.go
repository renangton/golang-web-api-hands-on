package usecase

import (
	"context"

	"dependency-injection-sample/domain/repository"
)

type Message interface {
	Get(ctx context.Context) string
}

type messageUseCase struct {
	messageRepo repository.Message
}

func NewMessage(messageRepo repository.Message) Message {
	return &messageUseCase{
		messageRepo: messageRepo,
	}
}

func (m *messageUseCase) Get(ctx context.Context) string {
	return m.messageRepo.RandomMessage()
}
