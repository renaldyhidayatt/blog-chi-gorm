package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type permissionService struct {
	permission dao.PermissionDao
}

func NewPermissionService(permission dao.PermissionDao) *permissionService {
	return &permissionService{permission: permission}
}

func (s *permissionService) FindPermission(IdUser int64, IdMenu int64) (entity.Permission, error) {
	res, err := s.permission.FindPermission(IdUser, IdMenu)

	return res, err
}

func (s *permissionService) CreatePermission(request []request.PermissionRequest) (bool, error) {
	res, err := s.permission.CreatePermission(request)

	return res, err
}

func (s *permissionService) UpdatePermission(request []request.PermissionRequest) (bool, error) {
	res, err := s.permission.UpdatePermission(request)

	return res, err
}

func (s *permissionService) DeletePermission(id int64) (bool, error) {
	res, err := s.permission.DeletePermission(id)

	return res, err

}
