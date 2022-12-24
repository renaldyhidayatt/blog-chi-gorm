package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type menuService struct {
	menu dao.MenuDao
}

func NewMenuService(menu dao.MenuDao) *menuService {
	return &menuService{menu: menu}
}

func (s *menuService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.menu.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *menuService) FindByMenu(id int64) (entity.Menu, error) {
	res, err := s.menu.FindByMenu(id)

	return res, err
}

func (s *menuService) CreateMenu(request request.MenuRequest) (bool, error) {
	res, err := s.menu.CreateMenu(request)

	return res, err
}

func (s *menuService) UpdateMenu(request request.MenuRequest, id int64) (bool, error) {
	res, err := s.menu.UpdateMenu(request, id)

	return res, err
}

func (s *menuService) DeleteMenu(id int64) (bool, error) {
	res, err := s.menu.DeleteMenu(id)

	return res, err
}
