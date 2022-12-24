package dao

import (
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
)

type PostDao interface {
	GetAll(*request.Pagination) (interface{}, error, int)
	CreatePost(request.PostRequest, string, []string, []string) (bool, error)
	FindPost(int64, int64) (entity.Post, error)
	UpdatePost(request.PostRequest, string, []string, []string) (bool, error)
	DeletePost(int64, int64) (bool, error)
	PublishPost(int64, int64) (bool, error)
	CancelPost(int64, int64) (bool, error)
}
