package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type authService struct {
	auth dao.AuthDao
}

func NewAuthService(auth dao.AuthDao) *authService {
	return &authService{auth: auth}
}

func (s *authService) CheckUser(loginRequest request.LoginRequest) (entity.User, error) {
	var schema request.LoginRequest

	schema.Username = loginRequest.Username
	schema.Password = loginRequest.Password

	res, err := s.auth.CheckUser(schema)

	return res, err
}

func (s *authService) ForgetPassword(forgetPassword request.ForgetPasswordRequest, userId int64) (entity.User, error) {
	var schema request.ForgetPasswordRequest

	schema.OldPassword = forgetPassword.OldPassword
	schema.NewPassword = forgetPassword.NewPassword

	res, err := s.auth.ForgetPassword(schema, userId)

	return res, err
}
