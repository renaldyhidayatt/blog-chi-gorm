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

type categoryHandler struct {
	category dao.CategoryDao
}

func NewCategoryHandler(category dao.CategoryDao) *categoryHandler {
	return &categoryHandler{category: category}
}

func (h *categoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.category.GetAll(pagination)

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

func (h *categoryHandler) FindCategory(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	get, err := h.category.FindCategory(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

func (h *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category request.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	category.Prepare()
	err = category.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.category.CreateCategory(category)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

func (h *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var category request.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	category.Prepare()
	err = category.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.category.UpdateCategory(category, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)

}

func (h *categoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	get, err := h.category.DeleteCategory(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
	}
}
