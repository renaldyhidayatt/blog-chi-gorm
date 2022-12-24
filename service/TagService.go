package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type tagService struct {
	tags dao.TagDao
}

func NewTagService(tags dao.TagDao) *tagService {
	return &tagService{tags: tags}
}

func (s *tagService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.tags.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *tagService) FindTag(id int64) (entity.Tag, error) {
	res, err := s.tags.FindTag(id)

	return res, err
}

func (s *tagService) CreateTag(request request.TagRequest) (bool, error) {
	res, err := s.tags.CreateTag(request)

	return res, err
}

func (s *tagService) UpdateTag(request request.TagRequest, id int64) (bool, error) {
	res, err := s.tags.UpdateTag(request, id)

	return res, err
}

func (s *tagService) DeleteTag(id int64) (bool, error) {
	res, err := s.tags.DeleteTag(id)

	return res, err
}
