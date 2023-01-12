package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewPermissionRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewPermissionRepository(db)
	service := service.NewPermissionService(repository)
	handler := handler.NewPermissionHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/{id_user:[0-9]+}/GetPermission/{id_menu:[0-9]+}", handler.FindPermission)
		r.Post("/CreatePermission", handler.CreatePermission)
		r.Put("/UpdatePermission", handler.UpdatePermission)
		r.Delete("/{id_permission:[0-9]+}/DeletePermission", handler.DeletePermission)
	})
}
