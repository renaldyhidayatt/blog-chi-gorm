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

type tagHandler struct {
	tag dao.TagDao
}

func NewTagHandler(tag dao.TagDao) *tagHandler {
	return &tagHandler{tag: tag}
}

// @Summary Get All Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /tag/GetAll [get]
func (h *tagHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.tag.GetAll(pagination)

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

// @Summary Find by Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /tag/{id_tag}/GetTag [get]
func (h *tagHandler) FindTag(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_tag"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.tag.FindTag(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

// @Summary Create Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param requestBody body request.TagRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /tag/CreateTag [post]
func (h *tagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tags request.TagRequest
	err := json.NewDecoder(r.Body).Decode(&tags)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	tags.Prepare()
	err = tags.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.tag.CreateTag(tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

// @Summary Update Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Param requestBody body request.TagRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /tag/{id_tag}/UpdateTag [put]
func (h *tagHandler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_tag"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var tag request.TagRequest
	err = json.NewDecoder(r.Body).Decode(&tag)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	tag.Prepare()
	err = tag.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.tag.UpdateTag(tag, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)

}

// @Summary Delete Tag
// @Description REST API Tag
// @Accept  json
// @Produce  json
// @Tags Tag Controller
// @Param id_tag path string true "Id Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /tag/{id_tag}/DeleteTag [delete]
func (h *tagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_tag"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.tag.DeleteTag(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)
	}
}
