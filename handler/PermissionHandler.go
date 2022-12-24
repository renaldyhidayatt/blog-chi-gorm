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
