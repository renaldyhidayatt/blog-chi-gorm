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

// @Summary Get All Article
// @Description REST API Article
// @Accept  json
// @Produce  json
// @Tags Article Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /article/GetAll [get]
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

// @Summary Find by Article
// @Description REST API Article
// @Accept  json
// @Produce  json
// @Tags Article Controller
// @Param id_article path string true "Id Article"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /article/{id_article}/GetArticle [get]
func (h *articleHandler) FindArticle(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.article.FindArticle(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

// @Summary Create Article
// @Description REST API Article
// @Accept  json
// @Produce  json
// @Tags Article Controller
// @Param requestBody body request.ArticleRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /article/CreateArticle [post]
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

// @Summary Update Article
// @Description REST API Article
// @Accept  json
// @Produce  json
// @Tags Article Controller
// @Param id_article path string true "Id Article"
// @Param requestBody body request.ArticleRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /article/{id_article}/UpdateArticle [put]
func (h *articleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	var article request.ArticleRequest
	err = json.NewDecoder(r.Body).Decode(&article)

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

// @Summary Delete Article
// @Description REST API Article
// @Accept  json
// @Produce  json
// @Tags Article Controller
// @Param id_article path string true "Id Article"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /article/{id_article}/DeleteArticle [delete]
func (h *articleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

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
