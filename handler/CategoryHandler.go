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

// @Summary Get All Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /category/GetAll [get]
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

// @Summary Find by Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id path string true "Id Category"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /category/{id} [get]
func (h *categoryHandler) FindCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.category.FindCategory(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

// @Summary Create Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param requestBody body request.CategoryRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /category/create [post]
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

	res, err := h.category.CreateCategory(category)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)

}

// @Summary Update Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id path string true "Id Category"
// @Param requestBody body request.CategoryRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /category/update/{id} [put]
func (h *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var category request.CategoryRequest

	err = json.NewDecoder(r.Body).Decode(&category)

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

	res, err := h.category.UpdateCategory(category, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", res, http.StatusOK)

}

// @Summary Delete Category
// @Description REST API Category
// @Accept  json
// @Produce  json
// @Tags Category Controller
// @Param id path string true "Id Category"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /category/delete/{id} [delete]
func (h *categoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.category.DeleteCategory(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)
	}
}
