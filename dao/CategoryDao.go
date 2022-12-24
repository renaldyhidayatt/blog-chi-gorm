package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type CategoryDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindCategory(int64) (entity.Category, error)
	CreateCategory(request.CategoryRequest) (bool, error)
	UpdateCategory(request.CategoryRequest, int64) (bool, error)
	DeleteCategory(int64) (bool, error)
}
