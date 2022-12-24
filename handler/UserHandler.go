package handler

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/payloads/request"
	"blog-chi-gorm/payloads/response"
	"blog-chi-gorm/security"
	"blog-chi-gorm/utils"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nfnt/resize"
)

type userHandler struct {
	user dao.UserDao
}

func NewUserHandler(user dao.UserDao) *userHandler {
	return &userHandler{user: user}
}

func (h *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.user.GetAll(pagination)

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

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userReq request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return

	}

	userReq.Prepare()
	err = userReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	hash, err := security.HashPassword(userReq.Password)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	userReq.Password = hash

	get, err := h.user.Insert(userReq)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", get, http.StatusCreated)
	}

}

func (h *userHandler) FindById(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	get, err := h.user.FindById(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mendapatkan data", get, http.StatusOK)
	}

}
func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var userReq request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	userReq.Prepare()
	err = userReq.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	hash, err := security.HashPassword(userReq.Password)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	userReq.Password = hash

	get, err := h.user.Update(userReq, Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil mengubah data", get, http.StatusOK)
	}

}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	get, err := h.user.Delete(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil menghapus data", get, http.StatusOK)
	}

}

func (h *userHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	find, err := h.user.FindById(Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Set Directory
	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	folderLocation := filepath.Join(dir, "images/profiles")

	if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
		os.MkdirAll(folderLocation, 0700)
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("photo")

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer file.Close()

	contentType := handler.Header.Get("Content-Type")

	if !CheckType(contentType) {
		response.ResponseError(w, http.StatusInternalServerError, errors.New("Format file tidak didukung"))
		return
	}

	var img image.Image

	switch contentType {
	case "image/jpg":
		img, err = jpeg.Decode(file)
	case "image/jpeg":
		img, err = jpeg.Decode(file)
	case "image/png":
		img, err = png.Decode(file)
	case "image/gif":
		img, err = gif.Decode(file)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	size := resize.Resize(300, 300, img, resize.Lanczos3)

	// Retrieve file information
	randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}

	filename := handler.Filename
	filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

	fileLocation := ""

	if find.Photo != "" {
		fileLocation = filepath.Join(dir, find.Photo)
	}

	exist, err := os.Stat(fileLocation)
	if exist != nil {
		e := os.Remove(fileLocation)
		if e != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

	out, err := os.Create(folderLocation + `/` + filename)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer out.Close()

	switch contentType {
	case "image/jpg":
		err = jpeg.Encode(out, size, nil)
	case "image/jpeg":
		err = jpeg.Encode(out, size, nil)
	case "image/png":
		err = png.Encode(out, size)
	case "image/gif":
		err = gif.Encode(out, size, nil)
	}

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := io.Copy(out, file); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	get, err := h.user.UploadImage(fmt.Sprintf("/images/profiles/%s", filename), Id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil upload image", get, http.StatusOK)

}

func CheckType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/jpg" || contentType == "image/gif"
}
