package repository

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type subMenuRepository struct {
	db *gorm.DB
}

func NewSubMenuRepository(db *gorm.DB) *subMenuRepository {
	return &subMenuRepository{db: db}
}

func (r *subMenuRepository) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var subMenus []entity.SubMenu

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

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

	find = find.Find(&subMenus)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = subMenus

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entity.SubMenu{}).Count(&counting).Error

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

func (r *subMenuRepository) FindBySubMenu(namaMenu string, namaSubMenu string) (entity.SubMenu, error) {
	var menu entity.Menu
	var subMenu entity.SubMenu

	r.db.Where("nama_menu=?", namaMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return entity.SubMenu{}, errors.New("nama menu tidak ditemukan")
	}

	r.db.Where("id_menu=? AND nama_sub_menu=?", menu.ID, namaSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return entity.SubMenu{}, errors.New("nama sub menu tidak ditemukan")
	}

	return subMenu, nil
}

func (r *subMenuRepository) CreateSubMenu(idMenu int64, subMenus []request.SubMenuRequest) (bool, error) {
	var menu entity.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.NamaMenu == "" {
		return false, errors.New("id menu tidak ditemukan")
	}

	for _, val := range subMenus {
		var subMenu entity.SubMenu

		r.db.Where("nama_sub_menu=?", val.NamaSubMenu).Take(&subMenu)

		if subMenu.NamaSubMenu != "" {
			return false, errors.New("nama sub menu sudah ada")
		}

		val.IdMenu = menu.ID
		val.Slug = slug.Make(val.NamaSubMenu)
		val.CreateAt = time.Now().Format("2006-01-02 15:04:05")
		val.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

		row := r.db.Omit("update_date").Create(&val)

		if row.Error != nil {
			return false, row.Error
		}
	}

	return true, nil
}

func (r *subMenuRepository) UpdateSubMenu(idMenu int64, idSubMenu int64, subMenus request.SubMenuRequest) (bool, error) {
	var menu entity.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("id Menu tidak ditemukan")
	}

	var subMenu entity.SubMenu

	r.db.Where("id_sub_menu=?", idSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return false, errors.New("id Sub Menu tidak ditemukan")
	}

	row := r.db.Omit("create_date").Model(&subMenus).Where("id_sub_menu=?", subMenu.ID).Updates(&subMenus)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *subMenuRepository) DeleteSubMenu(idMenu int64, idSubMenu int64) (bool, error) {
	var menu entity.Menu

	r.db.Where("id_menu=?", idMenu).Take(&menu)

	if menu.ID == 0 {
		return false, errors.New("id Menu tidak ditemukan")
	}

	var subMenu entity.SubMenu

	r.db.Where("id_sub_menu=?", idSubMenu).Take(&subMenu)

	if subMenu.NamaSubMenu == "" {
		return false, errors.New("id Sub Menu tidak ditemukan")
	}

	row := r.db.Model(&entity.SubMenu{}).Where("id_sub_menu=?", idSubMenu).Delete(subMenu)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}
