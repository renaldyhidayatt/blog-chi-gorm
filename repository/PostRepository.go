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

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAll(pagination *request.Pagination) (interface{}, error, int) {
	var posts []entity.Post

	totalRows := 0
	totalPages := 0
	fromRow := 0
	toRow := 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := r.db.Debug().Where("published=?", 1).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

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

	find = find.Find(&posts)

	if find.Error != nil {
		return nil, find.Error, totalPages
	}

	pagination.Rows = posts

	counting := int64(totalRows)

	// count all data
	err := r.db.Model(&entity.Post{}).Count(&counting).Error

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

func (r *postRepository) CreatePost(postRequest request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	var article entity.Article
	var post entity.Post

	r.db.Where("id_article=?", postRequest.IdArticle).Take(&article)

	if article.ID == 0 {
		return false, errors.New("article tidak ditemukan")
	}

	r.db.Where("nama_post=?", postRequest.NamaPost).Take(&post)

	if post.NamaPost != "" {
		return false, errors.New("title sudah ada")
	}

	for _, val := range categories {
		var category entity.Category
		r.db.Where("id_category=?", val).Take(&category)

		if category.ID == 0 {
			return false, errors.New("category tidak ditemukan")
		}
	}

	for _, val := range tags {
		var tag entity.Tag
		r.db.Where("id_tag=?", val).Take(&tag)

		if tag.ID == 0 {
			return false, errors.New("tag tidak ditemukan")
		}
	}

	err := r.db.Exec("INSERT INTO tb_posts (nama_post, slug, image, description, published, id_article, create_by, create_at) VALUES (?, ?, ?, ?, ?, ?, ? ,?)",
		postRequest.NamaPost,
		postRequest.Slug,
		urlImage,
		postRequest.Description,
		postRequest.Published,
		postRequest.IdArticle,
		postRequest.CreateBy,
		postRequest.CreateAt,
	).Error

	if err != nil {
		return false, err
	}

	var getPost entity.Post
	r.db.Where("nama_post=?", postRequest.NamaPost).Take(&getPost)

	if getPost.ID != 0 {
		for _, set := range categories {
			err := r.db.Exec("INSERT INTO tb_posts_has_categories (id_post, id_category) VALUES(?,?)", getPost.ID, set).Error

			if err != nil {
				return false, err
			}
		}

		for _, set := range tags {
			err := r.db.Exec("INSERT INTO tb_posts_has_tags (id_post, id_tag) VALUES(?,?)", getPost.ID, set).Error

			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func (r *postRepository) FindPost(IdArticle int64, IdPost int64) (entity.Post, error) {
	var post entity.Post

	r.db.Where("id_post=? AND id_article=?", IdPost, IdArticle).Take(&post)

	if post.ID <= 0 {
		return entity.Post{}, errors.New("post tidak ditemukan")
	}

	return post, nil
}

func (r *postRepository) UpdatePost(postRequest request.PostRequest, urlImage string, categories []string, tags []string) (bool, error) {
	var article entity.Article
	var post entity.Post

	r.db.Where("id_post=? AND id_article=?", postRequest.ID, postRequest.IdArticle).Take(&post)

	if post.ID == 0 {
		return false, errors.New("post tidak ditemukan")
	}

	r.db.Where("id_article=?", postRequest.IdArticle).Take(&article)

	if article.ID == 0 {
		return false, errors.New("article tidak ditemukan")
	}

	postCategory := r.db.Model(&entity.PostCategory{}).Where("id_post=?", post.ID).Delete(&entity.PostCategory{})

	if postCategory.Error != nil {
		return false, postCategory.Error
	}

	postTag := r.db.Model(&entity.PostTag{}).Where("id_post=?", post.ID).Delete(&entity.PostTag{})

	if postTag.Error != nil {
		return false, postTag.Error
	}

	err := r.db.Exec("UPDATE tb_posts SET nama_post=?, slug=?, image=?, description=?, published=?, id_article=?, update_by=?, update_at=? WHERE id_post=?",
		postRequest.NamaPost,
		postRequest.Slug,
		urlImage,
		postRequest.Description,
		postRequest.Published,
		article.ID,
		postRequest.UpdateBy,
		postRequest.UpdateAt,
		post.ID,
	).Error

	if err != nil {
		return false, err
	}

	var getPost entity.Post
	r.db.Where("nama_post=?", postRequest.NamaPost).Take(&getPost)

	if getPost.ID != 0 {
		for _, set := range categories {
			err := r.db.Exec("INSERT INTO tb_posts_has_categories (id_post, id_category) VALUES(?,?)", getPost.ID, set).Error

			if err != nil {
				return false, err
			}
		}

		for _, set := range tags {
			err := r.db.Exec("INSERT INTO tb_posts_has_tags (id_post, id_tag) VALUES(?,?)", getPost.ID, set).Error

			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func (r *postRepository) DeletePost(IdArticle int64, IdPost int64) (bool, error) {
	var post entity.Post

	r.db.Where("id_post=? AND id_article=?", IdPost, IdArticle).Take(&post)

	if post.ID <= 0 {
		return false, errors.New("post tidak ditemukan")
	}

	row := r.db.Model(&entity.Post{}).Where("id_post=?", post.ID).Delete(post)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *postRepository) PublishPost(IdArticle int64, IdPost int64) (bool, error) {
	var post entity.Post

	r.db.Where("id_post=? AND id_article=?", IdPost, IdArticle).Take(&post)

	if post.ID <= 0 {
		return false, errors.New("post tidak ditemukan")
	}

	row := r.db.Model(&entity.Post{}).Select("published").Where("id_post=?", post.ID).Update("published", true)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (r *postRepository) CancelPost(IdArticle int64, IdPost int64) (bool, error) {
	var post entity.Post

	r.db.Where("id_post=? AND id_article=?", IdPost, IdArticle).Take(&post)

	if post.ID <= 0 {
		return false, errors.New("post tidak ditemukan")
	}

	row := r.db.Model(&entity.Post{}).Select("published").Where("id_post=?", post.ID).Update("published", false)

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}
