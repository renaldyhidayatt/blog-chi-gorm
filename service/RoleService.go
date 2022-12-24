package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type roleService struct {
	role dao.RoleDao
}

func NewRoleService(role dao.RoleDao) *roleService {
	return &roleService{role: role}
}

func (s *roleService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.role.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *roleService) FindByNamaRole(name string) (entity.Role, error) {
	res, err := s.role.FindByNamaRole(name)

	return res, err
}

func (s *roleService) FindById(id int64) (entity.Role, error) {
	res, err := s.role.FindById(id)

	return res, err
}

func (s *roleService) Insert(request request.RoleRequest) (bool, error) {
	res, err := s.role.Insert(request)

	return res, err
}

func (s *roleService) Update(request request.RoleRequest, id int64) (bool, error) {
	res, err := s.role.Update(request, id)

	return res, err
}

func (s *roleService) Delete(id int64) (bool, error) {
	res, err := s.role.Delete(id)
	return res, err
}
