package gfa

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type Handler struct {
	App *firebase.App
	Auth *auth.Client
}

func (h *Handler) Init(app *firebase.App) error {
	var err error
	h.App = app
	h.Auth, err = app.Auth(context.Background())
	return err
}
