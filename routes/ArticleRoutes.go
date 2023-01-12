package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewArticleRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewArticleRepository(db)
	service := service.NewArticleService(repository)
	handler := handler.NewArticleHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Get("/{id_article:[0-9]+}/GetArticle", handler.FindArticle)
		r.Post("/CreateArticle", handler.CreateArticle)
		r.Put("/{id:[0-9]+}/UpdateArticle", handler.UpdateArticle)
		r.Delete("/{id_article:[0-9]+}/DeleteArticle", handler.DeleteArticle)
	})
}
