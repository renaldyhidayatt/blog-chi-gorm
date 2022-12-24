package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewMenuRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewMenuRepository(db)
	service := service.NewMenuService(repository)
	handler := handler.NewMenuHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Get("/{id_menu}/GetMenu", handler.FindByMenu)
		r.Post("/CreateMenu", handler.CreateMenu)
		r.Put("/{id_menu}/UpdateMenu", handler.UpdateMenu)
		r.Delete("/{id_menu}/DeleteMenu", handler.DeleteMenu)
	})
}
