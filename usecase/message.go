package usecase

import (
	"context"
	"math/rand"
)

type Message interface {
	Get(ctx context.Context) string
}

type messageUseCase struct{}

func NewMessage() Message {
	return &messageUseCase{}
}

func (m *messageUseCase) Get(ctx context.Context) string {
	message := []string{
		"There is always light behind the clouds.",
		"Change before you have to.",
		"If you can dream it, you can do it.",
		"Love the life you live. Live the life you love.",
	}
	return message[rand.Intn(len(message))]
}
