package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type SubMenuDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindBySubMenu(string, string) (entity.SubMenu, error)
	CreateSubMenu(int64, []request.SubMenuRequest) (bool, error)
	UpdateSubMenu(int64, int64, request.SubMenuRequest) (bool, error)
	DeleteSubMenu(int64, int64) (bool, error)
}
