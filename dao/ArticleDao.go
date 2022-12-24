package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type ArticleDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	FindArticle(int64) (entity.Article, error)
	CreateArticle(request.ArticleRequest) (bool, error)
	UpdateArticle(request.ArticleRequest, int64) (bool, error)
	DeleteArticle(int64) (bool, error)
}
