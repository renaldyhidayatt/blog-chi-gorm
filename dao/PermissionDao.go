package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type PermissionDao interface {
	FindPermission(int64, int64) (entity.Permission, error)
	CreatePermission([]request.PermissionRequest) (bool, error)
	UpdatePermission([]request.PermissionRequest) (bool, error)
	DeletePermission(int64) (bool, error)
}
