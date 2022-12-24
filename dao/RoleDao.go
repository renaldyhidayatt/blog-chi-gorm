package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type RoleDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindByNamaRole(string) (entity.Role, error)
	FindById(int64) (entity.Role, error)
	Insert(request.RoleRequest) (bool, error)
	Update(request.RoleRequest, int64) (bool, error)
	Delete(int64) (bool, error)
}
