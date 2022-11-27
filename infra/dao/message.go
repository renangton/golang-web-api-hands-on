package dao

import (
	"math/rand"

	"dependency-injection-sample/domain/repository"
)

type messageRepository struct{}

func NewMessage() repository.Message {
	return &messageRepository{}
}

func (m *messageRepository) RandomMessage() string {
	message := []string{
		"There is always light behind the clouds.",
		"Change before you have to.",
		"If you can dream it, you can do it.",
		"Love the life you live. Live the life you love.",
	}
	return message[rand.Intn(len(message))]
}
