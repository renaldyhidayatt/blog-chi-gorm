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

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var tags []entity.Tag

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

	find = find.Find(&tags)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = tags

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entity.Tag{}).Count(&counting).Error

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

func (r *tagRepository) FindTag(idTag int64) (entity.Tag, error) {
	var tag entity.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return entity.Tag{}, errors.New("Id tag tidak ditemukan")
	}

	return tag, nil
}

func (r *tagRepository) CreateTag(tagRequest request.TagRequest) (bool, error) {
	var tag entity.Tag

	r.db.Where("nama_tag=?", tagRequest.NamaTag).Take(&tag)

	if tag.ID > 0 {
		return false, errors.New("Nama tag sudah ada")
	}

	row := r.db.Omit("update_at").Create(&tagRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *tagRepository) UpdateTag(tagRequest request.TagRequest, idTag int64) (bool, error) {
	var tag entity.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return false, errors.New("Id tag tidak ditemukan")
	}

	row := r.db.Omit("create_at").Model(&tagRequest).Where("id_tag=?", tag.ID).Updates(&tagRequest)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *tagRepository) DeleteTag(idTag int64) (bool, error) {
	var tag entity.Tag

	r.db.Where("id_tag=?", idTag).Take(&tag)

	if tag.ID == 0 {
		return false, errors.New("Id tag tidak ditemukan")
	}

	row := r.db.Model(&entity.Tag{}).Where("id_tag=?", idTag).Delete(tag)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}
