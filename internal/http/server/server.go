package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pariip/notes-go/internal/contract"
	"net/http"
)

type (
	httpServer struct {
		handler *handler
		admin   *echo.Group
		user    *echo.Group
	}
)

var e = echo.New()

func NewHttpServer(h *handler) contract.HttpServer {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	admin := e.Group("/admin")
	user := e.Group("/user")
	return &httpServer{
		handler: h,
		admin:   admin,
		user:    user,
	}
}

func (s *httpServer) Start(port uint) error {
	s.setRoutes()
	if port == 0 {
		port = 8083
	}
	return e.Start(fmt.Sprintf(":%d", port))
}
