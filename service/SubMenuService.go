package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type submenuService struct {
	submenu dao.SubMenuDao
}

func NewSubMenuService(submenu dao.SubMenuDao) *submenuService {
	return &submenuService{submenu: submenu}
}

func (s *submenuService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.submenu.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *submenuService) FindBySubMenu(namaMenu string, namaSubMenu string) (entity.SubMenu, error) {
	res, err := s.submenu.FindBySubMenu(namaMenu, namaSubMenu)

	return res, err
}

func (s *submenuService) CreateSubMenu(idMenu int64, subMenus []request.SubMenuRequest) (bool, error) {
	res, err := s.submenu.CreateSubMenu(idMenu, subMenus)

	return res, err
}

func (s *submenuService) UpdateSubMenu(idMenu int64, idSubMenu int64, subMenus request.SubMenuRequest) (bool, error) {
	res, err := s.submenu.UpdateSubMenu(idMenu, idSubMenu, subMenus)

	return res, err
}

func (s *submenuService) DeleteSubMenu(idMenu int64, idSubMenu int64) (bool, error) {
	res, err := s.submenu.DeleteSubMenu(idMenu, idSubMenu)

	return res, err
}
