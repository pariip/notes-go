package server

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
	golang "log"
)

type (
	handler struct {
		cfg         *config.Config
		userService contract.UserService
		noteService contract.NoteService
		authService contract.AuthService
		logger      log.Logger
		translator  translate.Translator
	}
	HandlerFields struct {
		Cfg         *config.Config
		UserService contract.UserService
		NoteService contract.NoteService
		AuthService contract.AuthService
		Logger      log.Logger
		Translator  translate.Translator
	}
)

func NewHttpHandler(h *HandlerFields) *handler {
	if h.Cfg == nil {
		golang.Fatal("handler config is nil")
	}
	if h.UserService == nil {
		golang.Fatal("handler userService is nil")
	}
	if h.Logger == nil {
		golang.Fatal("handler logger is nil")
	}
	if h.Translator == nil {
		golang.Fatal("handler translator is nil")
	}
	return &handler{
		cfg:         h.Cfg,
		userService: h.UserService,
		noteService: h.NoteService,
		authService: h.AuthService,
		logger:      h.Logger,
		translator:  h.Translator,
	}
}
