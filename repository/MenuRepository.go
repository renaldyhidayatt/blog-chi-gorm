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

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var menus []entity.Menu

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Preload("SubMenu").Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s=?", column)
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

	find = find.Find(&menus)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = menus

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entity.Menu{}).Count(&counting).Error

	if err != nil {
		return nil, err, totalPages
	}

	totalRows = int(counting)

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
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
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil, totalPages

}

func (r *menuRepository) FindByMenu(menuId int64) (entity.Menu, error) {
	var menu entity.Menu

	r.db.Preload("SubMenu").Where("id_menu=?", menuId).Take(&menu)

	if menu.NamaMenu == "" {
		return entity.Menu{}, errors.New("nama menu tidak ditemukan")
	}

	return menu, nil

}

func (r *menuRepository) CreateMenu(menuRequest request.MenuRequest) (bool, error) {
	var menu entity.Menu

	r.db.Where("nama_menu=?", menuRequest.NamaMenu).Take(&menu)

	if menu.ID > 0 {
		return false, errors.New("nama menu sudah ada")
	}

	row := r.db.Omit("update_date").Create(&menuRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}

func (r *menuRepository) UpdateMenu(menuRequest request.MenuRequest, menuId int64) (bool, error) {
	var menu entity.Menu

	r.db.Where("id_menu=?", menuId).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("id Menu tidak ditemukan")
	}

	row := r.db.Omit("create_date").Model(&menuRequest).Where("id_menu=?", menu.ID).Updates(&menuRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}

func (r *menuRepository) DeleteMenu(menuId int64) (bool, error) {
	var menu entity.Menu

	r.db.Where("id_menu=?", menuId).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("Id menu tidak ditemukan")
	}

	row := r.db.Model(&entity.Menu{}).Where("id_menu=?", menuId).Delete(menu)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil

}
