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

type roleHandler struct {
	role dao.RoleDao
}

func NewRoleHandler(role dao.RoleDao) *roleHandler {
	return &roleHandler{role: role}
}

// @Summary Get All Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /role/GetAll [get]
func (h *roleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.role.GetAll(pagination)

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

// @Summary Find Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param role_name path string true "Role name"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /role/{role_name}/FindByRoleName [post]
func (h *roleHandler) FindByNamaRole(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "role_name")

	res, err := h.role.FindByNamaRole(name)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}
}

func (h *roleHandler) FindById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	res, err := h.role.FindById(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
	}

}

// @Summary Create Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param reqBody body request.RoleRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /role/CreateRole [post]
func (h *roleHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var roleReq request.RoleRequest
	err := json.NewDecoder(r.Body).Decode(&roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	roleReq.Prepare()
	err = roleReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.role.Insert(roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
	}

}

// @Summary Update Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param id_role path string true "Id Role"
// @Param reqBody body request.RoleRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /role/{id_role}/UpdateRole [put]
func (h *roleHandler) Update(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_role"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var roleReq request.RoleRequest
	err = json.NewDecoder(r.Body).Decode(&roleReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	roleReq.Prepare()
	err = roleReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	get, err := h.role.Update(roleReq, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
	}

}

// @Summary Delete Role
// @Description REST API Role
// @Accept  json
// @Produce  json
// @Tags Role Controller
// @Param id_role path string true "Id Role"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /role/{id_role}/DeleteRole [delete]
func (h *roleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id_role"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	res, err := h.role.Delete(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)
	}

}
