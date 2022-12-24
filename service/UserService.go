package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type userService struct {
	user dao.UserDao
}

func NewUserService(user dao.UserDao) *userService {
	return &userService{user: user}
}

func (s *userService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.user.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *userService) FindById(idUser int64) (entity.User, error) {
	res, err := s.user.FindById(idUser)

	return res, err
}

func (s *userService) Insert(request request.UserRequest) (bool, error) {
	res, err := s.user.Insert(request)

	return res, err
}

func (s *userService) Update(request request.UserRequest, idUser int64) (bool, error) {
	res, err := s.user.Update(request, idUser)

	return res, err
}

func (s *userService) Delete(idUser int64) (bool, error) {
	res, err := s.user.Delete(idUser)

	return res, err
}

func (s *userService) UploadImage(filename string, idUser int64) (bool, error) {
	res, err := s.user.UploadImage(filename, idUser)

	return res, err
}
