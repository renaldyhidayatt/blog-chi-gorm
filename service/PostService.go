package service

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/entity"
	"blog-chi-gorm/payloads/request"
	"errors"
)

type postService struct {
	posts dao.PostDao
}

func NewPostService(posts dao.PostDao) *postService {
	return &postService{posts: posts}
}

func (s *postService) GetAll(request *request.Pagination) (interface{}, error, int) {
	set, err, totalPages := s.posts.GetAll(request)

	if err != nil {
		return nil, errors.New("error"), 0
	}

	return set, err, totalPages
}

func (s *postService) FindPost(IdArticle int64, IdPost int64) (entity.Post, error) {
	res, err := s.posts.FindPost(IdArticle, IdPost)

	return res, err
}

func (s *postService) CreatePost(request request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	res, err := s.posts.CreatePost(request, urlImage, categories, tags)

	return res, err
}

func (s *postService) UpdatePost(postRequest request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	res, err := s.posts.UpdatePost(postRequest, urlImage, categories, tags)

	return res, err
}

func (s *postService) DeletePost(IdArticle int64, IdPost int64) (bool, error) {
	res, err := s.posts.DeletePost(IdArticle, IdPost)

	return res, err
}

func (s *postService) PublishPost(IdArticle int64, IdPost int64) (bool, error) {
	res, err := s.posts.PublishPost(IdArticle, IdPost)

	return res, err
}

func (s *postService) CancelPost(IdArticle int64, IdPost int64) (bool, error) {
	res, err := s.posts.CancelPost(IdArticle, IdPost)

	return res, err
}
