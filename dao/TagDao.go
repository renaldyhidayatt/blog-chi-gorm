package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type TagRepository interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindTag(int64) (entity.Tag, error)
	CreateTag(request.TagRequest) (bool, error)
	UpdateTag(request.TagRequest, int64) (bool, error)
	DeleteTag(int64) (bool, error)
}
