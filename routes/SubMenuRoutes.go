package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewSubMenuRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewSubMenuRepository(db)
	service := service.NewSubMenuService(repository)
	handler := handler.NewSubMenuHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/GetAll", handler.GetAll)
		r.Get("/{nama_menu}/GetSubMenu/{nama_sub_menu}", handler.FindBySubMenu)
		r.Post("/{id_menu:[0-9]+}/CreateSubMenu", handler.CreateSubMenu)
		r.Put("/{id_menu:[0-9]+}/UpdateSubMenu/{id_sub_menu:[0-9]+}", handler.UpdateSubMenu)
		r.Delete("/{id_menu:[0-9]+}/DeleteSubMenu/{id_sub_menu:[0-9]+}", handler.DeleteSubMenu)
	})
}
