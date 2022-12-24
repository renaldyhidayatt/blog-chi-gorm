package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type categoryService struct {
	category dao.CategoryDao
}

func NewCategoryService(category dao.CategoryDao) *categoryService {
	return &categoryService{category: category}
}

func (s *categoryService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.category.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *categoryService) FindCategory(id int64) (entity.Category, error) {
	res, err := s.category.FindCategory(id)

	return res, err
}

func (s *categoryService) CreateCategory(request request.CategoryRequest) (bool, error) {
	res, err := s.category.CreateCategory(request)

	return res, err
}

func (s *categoryService) UpdateCategory(request request.CategoryRequest, id int64) (bool, error) {
	res, err := s.category.UpdateCategory(request, id)

	return res, err
}

func (s *categoryService) DeleteCategory(id int64) (bool, error) {
	res, err := s.category.DeleteCategory(id)

	return res, err
}
