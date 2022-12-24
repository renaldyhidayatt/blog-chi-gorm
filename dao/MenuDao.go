package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type MenuDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindByMenu(int64) (entity.Menu, error)
	CreateMenu(request.MenuRequest) (bool, error)
	UpdateMenu(request.MenuRequest, int64) (bool, error)
	DeleteMenu(int64) (bool, error)
}
