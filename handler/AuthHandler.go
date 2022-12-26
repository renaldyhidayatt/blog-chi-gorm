package handler

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/payloads/request"
	"blog-chi-gorm/payloads/response"
	"blog-chi-gorm/security"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type authHandler struct {
	auth dao.AuthDao
}

func NewAuthHandler(auth dao.AuthDao) *authHandler {
	return &authHandler{auth: auth}
}

// @Summary Login
// @Description REST API Auth
// @Accept  json
// @Produce  json
// @Tags Auth Controller
// @Param reqBody body request.LoginRequest true "Form Request"
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /auth/login [post]
func (h *authHandler) CheckUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = loginRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return

	}

	get, err := h.auth.CheckUser(loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return

	}

	if get.ID > 0 {
		hashPwd := get.Password
		pwd := loginRequest.Password

		hash := security.VerifyPassword(hashPwd, pwd)

		if hash == nil {
			token, err := security.GenerateToken(get.Username)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, get, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password tidak sesuai"))
			return
		}

	}

}

// @Summary Forget Password
// @Description REST API Auth
// @Accept  json
// @Produce  json
// @Tags Auth Controller
// @Param id path string true "Id User"
// @Param reqBody body request.ForgetPasswordRequest true "Form Request"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /auth/{id}/ForgetPassword [post]
func (h *authHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var forgetRequest request.ForgetPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&forgetRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	get, err := h.auth.ForgetPassword(forgetRequest, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if get.ID > 0 {
		hashPwd := get.Password
		newPwd := forgetRequest.NewPassword
		oldPwd := forgetRequest.OldPassword

		err = security.VerifyPassword(hashPwd, oldPwd)

		if err == nil {
			if newPwd != oldPwd {
				hash, err := security.HashPassword(newPwd)

				if err != nil {
					response.ResponseError(w, http.StatusInternalServerError, err)
					return
				}

				forgetRequest.NewPassword = hash

				set, err := h.auth.ForgetPassword(forgetRequest, Id)

				if err != nil {
					response.ResponseError(w, http.StatusInternalServerError, err)
					return
				}

				response.ResponseMessage(w, "Berhasil mengubah password", set, http.StatusOK)
			} else {
				response.ResponseError(w, http.StatusBadRequest, errors.New("password baru tidak boleh sama"))
				return
			}
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password lama tidak sesuai"))
			return
		}
	}

}
