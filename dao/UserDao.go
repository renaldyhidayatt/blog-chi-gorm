package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type UserDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindById(int64) (entity.User, error)
	Insert(request.UserRequest) (bool, error)
	Update(request.UserRequest, int64) (bool, error)
	Delete(int64) (bool, error)

	UploadImage(string, int64) (bool, error)
}
