package handler

import (
	"math/rand"
	"net/http"

	"dependency-injection-sample/usecase"

	"github.com/go-chi/render"
)

type Message interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct {
	useCase usecase.Message
}

func NewMessage(useCase usecase.Message) Message {
	return &messageHandler{
		useCase: useCase,
	}
}

func (m *messageHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": m.useCase.Get(ctx),
	})
}
