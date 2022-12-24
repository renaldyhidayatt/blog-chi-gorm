package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewCategoryRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewCategoryRepository(db)
	service := service.NewCategoryService(repository)
	handler := handler.NewCategoryHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Post("/create", handler.CreateCategory)
		r.Get("/{id}", handler.FindCategory)
		r.Put("/update/{id}", handler.UpdateCategory)
		r.Delete("/delete/{id}", handler.Delete)

	})
}
