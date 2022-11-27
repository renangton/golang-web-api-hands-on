package handler

import (
	"math/rand"
	"net/http"

	"github.com/go-chi/render"
)

type Message interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct{}

func NewMessage() Message {
	return &messageHandler{}
}

func (m *messageHandler) Get(w http.ResponseWriter, r *http.Request) {
	message := []string{
		"There is always light behind the clouds.",
		"Change before you have to.",
		"If you can dream it, you can do it.",
		"Love the life you live. Live the life you love.",
	}
	randomNumber := rand.Intn(len(message))
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": message[randomNumber],
	})
}
