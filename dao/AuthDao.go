package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type AuthDao interface {
	CheckUser(request.LoginRequest) (entity.User, error)
	ForgetPassword(request.ForgetPasswordRequest, int64) (entity.User, error)
}
