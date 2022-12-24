package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type articleService struct {
	article dao.ArticleDao
}

func NewArticleService(article dao.ArticleDao) *articleService {
	return &articleService{article: article}
}

func (s *articleService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.article.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *articleService) FindArticle(id int64) (entity.Article, error) {
	res, err := s.article.FindArticle(id)

	return res, err
}

func (s *articleService) CreateArticle(request request.ArticleRequest) (bool, error) {
	res, err := s.article.CreateArticle(request)

	return res, err
}

func (s *articleService) UpdateArticle(request request.ArticleRequest, id int64) (bool, error) {
	res, err := s.article.UpdateArticle(request, id)

	return res, err
}

func (s *articleService) DeleteArticle(id int64) (bool, error) {
	res, err := s.article.DeleteArticle(id)

	return res, err
}
