package repository

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var user []entity.User

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	searchs := pagination.Searchs

	if searchs != nil {
		for _, val := range searchs {
			column := val.Column
			action := val.Action
			query := val.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			}
		}
	}

	find = find.Find(&user)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = user

	counting := int64(totalRows)

	// count all
	err := r.db.Model(&entity.User{}).Count(&counting).Error

	if err != nil {
		return nil, find.Error, totalPages
	}

	totalRows = int(counting)
	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil, totalPages

}

func (r *userRepository) FindById(idUser int64) (entity.User, error) {
	var user entity.User

	r.db.Where("id_user=?", idUser).Preload("Role").Take(&user)

	if user.ID <= 0 {
		return entity.User{}, errors.New("id User tidak ditemukan")
	}

	return user, nil

}

func (r *userRepository) Insert(reqUser request.UserRequest) (bool, error) {
	var user entity.User

	r.db.Where("email=? OR username=?", reqUser.Email, reqUser.Username).Take(&user)

	if user.Username != "" {
		return false, errors.New("username sudah terdaftar")
	} else if user.Email != "" {
		return false, errors.New("email sudah terdaftar")
	} else if user.Username != "" && user.Email != "" {
		return false, errors.New("username dan Email sudah terdaftar")
	}

	err := r.db.Exec("INSERT INTO tb_users (first_name, last_name, username, password, email, no_telp, photo, id_role, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		reqUser.FirstName,
		reqUser.LastName,
		reqUser.Username,
		reqUser.Password,
		reqUser.Email,
		reqUser.NoTelp,
		reqUser.Photo,
		reqUser.IdRole,
		reqUser.Status,
	).Error

	if err != nil {
		return false, err
	}

	return true, nil

}

func (r *userRepository) Update(reqUser request.UserRequest, idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Omit("create_at, photo, username").Model(&reqUser).Where("id_user=?", get.ID).Updates(&reqUser)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}

func (r *userRepository) Delete(idUser int64) (bool, error) {
	get, err := r.FindById(idUser)

	if err != nil {
		return false, err
	}

	row := r.db.Model(&entity.User{}).Where("id_user=?", get.ID).Delete(get)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}

func (r *userRepository) UploadImage(filename string, userId int64) (bool, error) {
	get, err := r.FindById(userId)

	if err != nil {
		return false, err
	}

	get.Photo = filename

	row := r.db.Model(&get).Select("photo").Where("id_user=?", get.ID).Updates(&get)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}
