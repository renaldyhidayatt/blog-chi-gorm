package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewPostRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewPostRepository(db)
	service := service.NewPostService(repository)
	handler := handler.NewPostHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Post("/create", handler.CreatePost)
		r.Get("/{id_article:[0-9]+}/FindPost/{id_post:[0-9]+}", handler.FindPost)
		r.Put("/update", handler.UpdatePost)
		r.Delete("/{id_article}/delete/{id_post:[0-9]+}", handler.DeletePost)
		r.Post("/{id_article}/PublishPost/{id_post:[0-9]+}", handler.PublishPost)
		r.Post("/{id_article}/CancelPost/{id_post:[0-9]+}", handler.CancelPost)
	})
}
