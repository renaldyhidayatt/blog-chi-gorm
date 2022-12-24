package handler

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/payloads/request"
	"blog-chi-gorm/payloads/response"
	"blog-chi-gorm/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type articleHandler struct {
	article dao.ArticleDao
}

func NewArticleHandler(article dao.ArticleDao) *articleHandler {
	return &articleHandler{article: article}
}

func (h *articleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.article.GetAll(pagination)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	err = utils.SetupPagination(r, set, pagination, totalPages)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", set, http.StatusOK)
	}
}

func (h *articleHandler) FindArticle(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	res, err := h.article.FindArticle(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *articleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var articles request.ArticleRequest
	err := json.NewDecoder(r.Body).Decode(&articles)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	articles.Prepare()
	err = articles.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.article.CreateArticle(articles)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

func (h *articleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	var article request.ArticleRequest
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	article.Prepare()
	err = article.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.article.UpdateArticle(article, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)

}

func (h *articleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := h.article.DeleteArticle(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)

}
