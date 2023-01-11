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

type menuHandler struct {
	menu dao.MenuDao
}

func NewMenuHandler(menu dao.MenuDao) *menuHandler {
	return &menuHandler{menu: menu}
}

// @Summary Get All Menu
// @Description REST API Menu
// @Accept  json
// @Produce  json
// @Tags Menu Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /menu/GetAll [get]
func (h *menuHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.menu.GetAll(pagination)

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

// @Summary Find by Menu
// @Description REST API Menu
// @Accept  json
// @Produce  json
// @Tags Menu Controller
// @Param id_menu path string true "Id Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /menu/{id_menu}/GetMenu [get]
func (h *menuHandler) FindByMenu(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.menu.FindByMenu(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}
}

// @Summary Create Menu
// @Description REST API Menu
// @Accept  json
// @Produce  json
// @Tags Menu Controller
// @Param requestBody body request.MenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /menu/CreateMenu [post]
func (h *menuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menus request.MenuRequest
	err := json.NewDecoder(r.Body).Decode(&menus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	menus.Prepare()
	err = menus.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.menu.CreateMenu(menus)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

// @Summary Update Menu
// @Description REST API Menu
// @Accept  json
// @Produce  json
// @Tags Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param requestBody body request.MenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /menu/{id_menu}/UpdateMenu [put]
func (h *menuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var menus request.MenuRequest
	err = json.NewDecoder(r.Body).Decode(&menus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	menus.Prepare()
	err = menus.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.menu.UpdateMenu(menus, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)

}

// @Summary Delete Menu
// @Description REST API Menu
// @Accept  json
// @Produce  json
// @Tags Menu Controller
// @Param id_menu path string true "Id Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /menu/{id_menu}/DeleteMenu [delete]
func (h *menuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.menu.DeleteMenu(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)

}
