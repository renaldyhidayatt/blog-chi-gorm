package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewTagRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewTagRepository(db)
	service := service.NewTagService(repository)
	handler := handler.NewTagHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/GetAll", handler.GetAll)
		r.Get("/{id_tag}/GetTag", handler.FindTag)
		r.Post("/CreateTag", handler.CreateTag)
		r.Put("/{id_tag}/UpdateTag", handler.UpdateTag)
		r.Delete("/{id_tag}/DeleteTag", handler.Delete)
	})
}
