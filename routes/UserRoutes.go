package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewUserRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Post("/create", handler.Create)
		r.Get("/{id:[0-9]+}", handler.FindById)
		r.Put("/update/{id:[0-9]+}", handler.Update)
		r.Delete("/delete/{id:[0-9]+}", handler.Delete)
		r.Post("/{id:[0-9]+}/UploadImage", handler.UploadImage)

	})
}
