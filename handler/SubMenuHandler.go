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

type subMenuHandler struct {
	submenu dao.SubMenuDao
}

func NewSubMenuHandler(submenu dao.SubMenuDao) *subMenuHandler {
	return &subMenuHandler{submenu: submenu}
}

// @Summary Get All Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /submenu/GetAll [get]
func (h *subMenuHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.submenu.GetAll(pagination)

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

// @Summary Find by Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param nama_menu path string true "Nama Menu"
// @Param nama_sub_menu path string true "Nama Sub Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /submenu/{nama_menu}/GetSubMenu/{nama_sub_menu} [get]
func (h *subMenuHandler) FindBySubMenu(w http.ResponseWriter, r *http.Request) {
	nama_menu := chi.URLParam(r, "nama_menu")
	nama_submenu := chi.URLParam(r, "nama_submenu")

	res, err := h.submenu.FindBySubMenu(nama_menu, nama_submenu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)

}

// @Summary Create Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param requestBody body []request.SubMenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /submenu/{id_menu}/CreateSubMenu [post]
func (h *subMenuHandler) CreateSubMenu(w http.ResponseWriter, r *http.Request) {
	id_menu, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var subMenus []request.SubMenuRequest
	err = json.NewDecoder(r.Body).Decode(&subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	for _, val := range subMenus {
		err = val.Validate()

		if err != nil {
			response.ResponseError(w, http.StatusBadRequest, err)
			return
		}
	}

	get, err := h.submenu.CreateSubMenu(id_menu, subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

// @Summary Update Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param id_sub_menu path string true "Id Sub Menu"
// @Param requestBody body request.SubMenuRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /submenu/{id_menu}/UpdateSubMenu/{id_sub_menu} [put]
func (h *subMenuHandler) UpdateSubMenu(w http.ResponseWriter, r *http.Request) {
	id_menu, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	id_sub_menu, err := strconv.ParseInt(chi.URLParam(r, "id_sub_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var subMenus request.SubMenuRequest
	err = json.NewDecoder(r.Body).Decode(&subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	subMenus.Prepare()
	err = subMenus.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	subMenus.IdMenu = id_menu

	res, err := h.submenu.UpdateSubMenu(id_menu, id_sub_menu, subMenus)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", res, http.StatusOK)

}

// @Summary Delete Sub Menu
// @Description REST API Sub Menu
// @Accept  json
// @Produce  json
// @Tags Sub Menu Controller
// @Param id_menu path string true "Id Menu"
// @Param id_sub_menu path string true "Id Sub Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /submenu/{id_menu}/DeleteSubMenu/{id_sub_menu} [delete]
func (h *subMenuHandler) DeleteSubMenu(w http.ResponseWriter, r *http.Request) {
	id_menu, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	id_sub_menu, err := strconv.ParseInt(chi.URLParam(r, "id_sub_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.submenu.DeleteSubMenu(id_menu, id_sub_menu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)

}
