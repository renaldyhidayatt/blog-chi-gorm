package handler

import (
	"blog-chi-gorm/dao"
	"blog-chi-gorm/payloads/request"
	"blog-chi-gorm/payloads/response"
	"blog-chi-gorm/utils"
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
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gosimple/slug"
	"github.com/nfnt/resize"
)

type postHandler struct {
	posts dao.PostDao
}

func NewPostHandler(posts dao.PostDao) *postHandler {
	return &postHandler{posts: posts}
}

// @Summary Get All Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param sort query string false "Sort"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/GetAll [get]
func (h *postHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := utils.SortPagination(r)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	}

	set, err, totalPages := h.posts.GetAll(pagination)

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

// @Summary Create Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param requestBody formData request.PostRequest true "Form"
// @Param photo formData file true "Photo"
// @Param categories formData string true "Category"
// @Param tags formData string true "Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/create [post]
func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	folderLocation := filepath.Join(dir, "images/posts")

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
		response.ResponseError(w, http.StatusInternalServerError, errors.New("format file tidak didukung"))
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

	size := resize.Resize(600, 600, img, resize.Lanczos3)

	// Retrieve file information
	randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}

	filename := handler.Filename
	filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

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

	publish, _ := strconv.ParseBool(r.FormValue("published"))
	idArticle, _ := strconv.ParseInt(r.FormValue("id_article"), 10, 64)

	posts := request.PostRequest{
		NamaPost:    r.FormValue("nama_post"),
		Description: r.FormValue("description"),
		Published:   publish,
		IdArticle:   idArticle,
		CreateBy:    r.FormValue("create_by"),
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateBy:    r.FormValue("update_by"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	urlImage := fmt.Sprintf("/images/posts/%s", filename)
	posts.Slug = slug.Make(posts.NamaPost)

	categories := strings.Split(r.FormValue("categories"), ",")
	tags := strings.Split(r.FormValue("tags"), ",")

	res, err := h.posts.CreatePost(posts, urlImage, categories, tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)

}

// @Summary Find Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/{id_article}/FindPost/{id_post} [get]
func (h *postHandler) FindPost(w http.ResponseWriter, r *http.Request) {
	id_article, _ := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)
	id_post, _ := strconv.ParseInt(chi.URLParam(r, "id_post"), 10, 64)

	res, err := h.posts.FindPost(id_article, id_post)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)

}

// @Summary Update Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param requestBody formData request.PostRequest true "Form"
// @Param photo formData file false "Photo"
// @Param categories formData string true "Category"
// @Param tags formData string true "Tag"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/update [put]
func (h *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	publish, _ := strconv.ParseBool(r.FormValue("published"))
	idPost, _ := strconv.ParseInt(r.FormValue("id_post"), 10, 64)
	idArticle, _ := strconv.ParseInt(r.FormValue("id_article"), 10, 64)

	posts := request.PostRequest{
		ID:          idPost,
		NamaPost:    r.FormValue("nama_post"),
		Description: r.FormValue("description"),
		Published:   publish,
		IdArticle:   idArticle,
		CreateBy:    r.FormValue("create_by"),
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateBy:    r.FormValue("update_by"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	find, err := h.posts.FindPost(idArticle, idPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	file, handler, err := r.FormFile("photo")
	urlImage := ""

	if err == nil {

		dir, err := os.Getwd()
		folderLocation := filepath.Join(dir, "images/posts")

		if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
			os.MkdirAll(folderLocation, 0700)
		}

		r.ParseMultipartForm(10 << 20)

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

		size := resize.Resize(600, 600, img, resize.Lanczos3)

		// Retrieve file information
		randomString := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, 10)

		for i := range b {
			b[i] = randomString[rand.Intn(len(randomString))]
		}

		filename := handler.Filename
		filename = fmt.Sprintf("%s%s", string(b), filepath.Ext(handler.Filename))

		fileLocation := ""

		if find.Image != "" {
			fileLocation = filepath.Join(dir, find.Image)
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

		urlImage = fmt.Sprintf("/images/posts/%s", filename)
	}

	fmt.Println(urlImage)

	if urlImage == "" {
		fmt.Println(find.Image)
		urlImage = find.Image
	}

	posts.Slug = slug.Make(posts.NamaPost)

	categories := strings.Split(r.FormValue("categories"), ",")
	tags := strings.Split(r.FormValue("tags"), ",")

	res, err := h.posts.UpdatePost(posts, urlImage, categories, tags)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mengubah data", res, http.StatusOK)

}

// @Summary Delete Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/{id_article}/DeletePost/{id_post} [delete]
func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	IdArticle, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	IdPost, err := strconv.ParseInt(chi.URLParam(r, "id_post"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	find, err := h.posts.FindPost(IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	dir, err := os.Getwd()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	fileLocation := ""

	if find.Image != "" {
		fileLocation = filepath.Join(dir, find.Image)
	}

	exist, err := os.Stat(fileLocation)
	if exist != nil {
		e := os.Remove(fileLocation)
		if e != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

	res, err := h.posts.DeletePost(IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil menghapus data", res, http.StatusOK)

}

// @Summary Publish Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/{id_article}/PublishPost/{id_post} [post]
func (h *postHandler) PublishPost(w http.ResponseWriter, r *http.Request) {
	IdArticle, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	IdPost, err := strconv.ParseInt(chi.URLParam(r, "id_post"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.posts.PublishPost(IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil mempublish data", res, http.StatusOK)

}

// @Summary Cancel Post
// @Description REST API Post
// @Accept  json
// @Produce  json
// @Tags Post Controller
// @Param id_article path string true "Id Article"
// @Param id_post path string true "Id Post"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Success 201 {object} response.Response
// @Failure 500,400,404,403 {object} response.Response
// @Router /posts/{id_article}/CancelPost/{id_post} [post]
func (h *postHandler) CancelPost(w http.ResponseWriter, r *http.Request) {
	IdArticle, err := strconv.ParseInt(chi.URLParam(r, "id_article"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	IdPost, err := strconv.ParseInt(chi.URLParam(r, "id_post"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.posts.CancelPost(IdArticle, IdPost)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseMessage(w, "Berhasil membatalkan data", res, http.StatusOK)

}
