package routes

import (
	"blog-chi-gorm/handler"
	"blog-chi-gorm/middlewares"
	"blog-chi-gorm/repository"
	"blog-chi-gorm/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewRoleRoutes(prefix string, db *gorm.DB, router *chi.Mux) {
	repository := repository.NewRoleRepository(db)
	service := service.NewRoleService(repository)
	handler := handler.NewRoleHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)

		r.Get("/GetAll", handler.GetAll)
		r.Post("/{role_name}/FindByRoleName", handler.FindByNamaRole)
		r.Post("/CreateRole", handler.Insert)
		r.Put("/{id_role:[0-9]+} ubah parameter ini jadi role_name/UpdateRole", handler.Update)
		r.Delete("/{id_role}/DeleteRole", handler.Delete)
	})
}
