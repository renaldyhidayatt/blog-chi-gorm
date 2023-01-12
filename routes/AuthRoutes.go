package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewAuthRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewAuthRepository(db)
	serviceAuth := service.NewAuthService(repository)
	handlerAuth := handler.NewAuthHandler(serviceAuth)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Post("/DoLogin", handlerAuth.CheckUser)
		r.Post("/{id:[0-9]+}/ForgotPassword", handlerAuth.ForgotPassword)
	})
}
