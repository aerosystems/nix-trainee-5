package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/go-chi/chi/v5"
)

func (h *BaseHandler) ReadPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postRepo.FindAll()
	if err != nil {
		log.Println(err)
	}

	if len(*posts) == 0 {
		_ = writeError(w, r, errors.New("posts do not exist"), http.StatusNotFound)
		return
	}

	payload := Response{
		Error:   false,
		Message: "all posts with ID were found successfully",
		Data:    posts,
	}
	_ = writeResponse(w, r, http.StatusOK, payload)
}

func (h *BaseHandler) ReadPost(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		log.Println(err)
	}

	if *post == (models.Post{}) {
		_ = writeError(w, r, errors.New("post with ID "+stringID+" does not exist"), http.StatusNotFound)
		return
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was found successfully",
		Data:    post,
	}
	_ = writeResponse(w, r, http.StatusOK, payload)
}

func (h *BaseHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
	}

	var requestPayload struct {
		UserID int    `json:"userId" xml:"userId"`
		Title  string `json:"title" xml:"title"`
		Body   string `json:"body" xml:"body"`
	}

	err = readRequest(w, r, &requestPayload)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
		return
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		log.Println(err)
	}

	if *post != (models.Post{}) {
		_ = writeError(w, r, errors.New("post with ID "+stringID+" exists"), http.StatusNotFound)
		return
	}

	newPost := models.Post{
		ID:     ID,
		UserID: requestPayload.UserID,
		Title:  requestPayload.Title,
		Body:   requestPayload.Body,
	}

	err = h.postRepo.Create(&newPost)
	if err != nil {
		log.Println(err)
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was created successfully",
		Data:    newPost,
	}
	_ = writeResponse(w, r, http.StatusOK, payload)
}

func (h *BaseHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
	}

	var requestPayload struct {
		UserID int    `json:"userId" xml:"userId"`
		Title  string `json:"title" xml:"title"`
		Body   string `json:"body" xml:"body"`
	}

	err = readRequest(w, r, &requestPayload)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
		return
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		log.Println(err)
	}

	if *post == (models.Post{}) {
		_ = writeError(w, r, errors.New("post with ID "+stringID+"does not exists"), http.StatusNotFound)
		return
	}

	newPost := post
	newPost.ID = ID

	if requestPayload.UserID != 0 {
		newPost.UserID = requestPayload.UserID
	}
	if requestPayload.Title != "" {
		newPost.Title = requestPayload.Title
	}
	if requestPayload.Body != "" {
		newPost.Body = requestPayload.Body
	}

	err = h.postRepo.Update(newPost)
	if err != nil {
		log.Println(err)
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was updated successfully",
		Data:    nil,
	}
	_ = writeResponse(w, r, http.StatusOK, payload)
}

func (h *BaseHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		_ = writeError(w, r, err, http.StatusBadRequest)
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		log.Println(err)
	}

	if *post == (models.Post{}) {
		_ = writeError(w, r, errors.New("post with ID "+stringID+" does not exist"), http.StatusNotFound)
		return
	}

	err = h.postRepo.Delete(post)
	if err != nil {
		log.Println(err)
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was deleted successfully",
		Data:    nil,
	}
	_ = writeResponse(w, r, http.StatusOK, payload)
}
