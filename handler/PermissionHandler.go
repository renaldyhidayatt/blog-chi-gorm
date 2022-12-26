package handler

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/payloads/request"
	"blog-chi-gorm/payloads/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type permissionHandler struct {
	permission dao.PermissionDao
}

func NewPermissionHandler(permission dao.PermissionDao) *permissionHandler {
	return &permissionHandler{permission: permission}
}

// @Summary Find Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param id_user path string true "Id User"
// @Param id_menu path string true "Id Menu"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /permission/{id_user}/GetPermission/{id_menu} [get]
func (h *permissionHandler) FindPermission(w http.ResponseWriter, r *http.Request) {
	id_user, err := strconv.ParseInt(chi.URLParam(r, "id_user"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	id_menu, err := strconv.ParseInt(chi.URLParam(r, "id_menu"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.permission.FindPermission(id_user, id_menu)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)

}

// @Summary Create Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param requestBody body []request.PermissionRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /permission/CreatePermission [post]
func (h *permissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission []request.PermissionRequest
	err := json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := h.permission.CreatePermission(permission)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)

}

// @Summary Update Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param requestBody body []request.PermissionRequest true "Form"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /permission/UpdatePermission [put]
func (h *permissionHandler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	var permission []request.PermissionRequest
	err := json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := h.permission.UpdatePermission(permission)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)

}

// @Summary Delete Permission
// @Description REST API Permission
// @Accept  json
// @Produce  json
// @Tags Permission Controller
// @Param id_permission path string true "Id Permission"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /permission/{id_permission}/DeletePermission [delete]
func (h *permissionHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_permission"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.permission.DeletePermission(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)
	}
}
