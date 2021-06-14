package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"net/http"
)

type (
	httpServer struct {
		handler *handler
		public  *echo.Group
		admin   *echo.Group
		user    *echo.Group
		note    *echo.Group
		image   *echo.Group
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
	e.Use(middleware.Recover())

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(h.cfg.Auth.JWTSecret),
		Claims:     &models.Claims{},
		ErrorHandler: func(err error) error {
			fmt.Println(err)

			return &echo.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			}
		},
	}
	public := e.Group("")
	admin := e.Group("/admin", middleware.JWTWithConfig(jwtConfig), middlewarePermission(h, types.Admin))
	user := e.Group("/user", middleware.JWTWithConfig(jwtConfig))
	note := e.Group("/note", middleware.JWTWithConfig(jwtConfig))
	image := e.Group("/image", middleware.JWTWithConfig(jwtConfig))

	return &httpServer{
		handler: h,
		public:  public,
		admin:   admin,
		user:    user,
		note:    note,
		image:   image,
	}
}

func (s *httpServer) Start(port uint) error {
	s.setRoutes()
	if port == 0 {
		port = 8083
	}
	return e.Start(fmt.Sprintf(":%d", port))
}
