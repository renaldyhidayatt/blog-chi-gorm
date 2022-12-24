package repository

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
	"time"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *permissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) FindPermission(IdUser int64, IdMenu int64) (entity.Permission, error) {
	var user entity.User

	r.db.Where("id_user=?", IdUser).Take(&user)

	if user.Username == "" {
		return entity.Permission{}, errors.New("Id User tidak ditemukan")
	}

	var menu entity.Menu

	r.db.Where("id_menu=?", IdMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return entity.Permission{}, errors.New("Id Menu tidak ditemukan")
	}

	var permission entity.Permission

	r.db.Preload("Menu").Preload("User").Where("id_user=? AND id_menu=?", user.ID, menu.ID).Take(&permission)

	if permission.ID == 0 || permission.ID < 1 {
		return entity.Permission{}, errors.New("Data tidak ditemukan")
	}

	return permission, nil
}

func (r *permissionRepository) CreatePermission(permission []request.PermissionRequest) (bool, error) {
	for _, val := range permission {
		var user entity.User

		r.db.Where("id_user=?", val.IdUser).Take(&user)

		if user.ID == 0 || user.ID < 1 {
			return false, errors.New("Id User tidak ditemukan")
		}

		var menu entity.Menu

		r.db.Where("id_menu=?", val.IdMenu).Take(&menu)

		if menu.ID == 0 {
			return false, errors.New("Id Menu tidak ditemukan")
		}

		var permis entity.Permission

		r.db.Where("id_user=? AND id_menu=?", user.ID, menu.ID).Take(&permis)

		if permis.ID != 0 {
			return false, errors.New("data sudah terdaftar")
		}

		val.CreateAt = time.Now().Format("2006-01-02 15:04:05")
		row := r.db.Create(&val)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *permissionRepository) UpdatePermission(permission []request.PermissionRequest) (bool, error) {
	for _, val := range permission {
		var permis entity.Permission

		r.db.Where("id_permission=?", val.ID).Take(&permis)

		if permis.ID == 0 {
			return false, errors.New("Id permission tidak ditemukan")
		}

		var user entity.User

		r.db.Where("id_user=?", val.IdUser).Take(&user)

		if user.ID == 0 || user.ID < 1 {
			return false, errors.New("Id User tidak ditemukan")
		}

		var menu entity.Menu

		r.db.Where("id_menu=?", val.IdMenu).Take(&menu)

		if menu.ID == 0 {
			return false, errors.New("Id Menu tidak ditemukan")
		}

		row := r.db.Exec("UPDATE tb_has_permission SET f_create=?, f_read=?, f_update=?, f_delete=?, f_publish=? WHERE id_permission=? AND id_menu=? AND id_user=?", val.FCreate, val.FRead, val.FUpdate, val.FDelete, val.FPublish, val.ID, val.IdMenu, val.IdUser)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *permissionRepository) DeletePermission(IdPermission int64) (bool, error) {
	var permission entity.Permission

	r.db.Where("id_permission=?", IdPermission).Take(&permission)

	if permission.ID == 0 {
		return false, errors.New("id Permission tidak ditemukan")
	}

	row := r.db.Model(&entity.Permission{}).Where("id_permission=?", IdPermission).Delete(permission)

	if row.Error != nil {
		return false, row.Error
	}
	return true, nil
}
