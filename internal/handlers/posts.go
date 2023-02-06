package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

// ReadPosts godoc
// @Summary get all posts
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Success 200 {object} Response{data=[]models.Post}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /posts [get]
func (h *BaseHandler) ReadPosts(c echo.Context) error {
	posts, err := h.postRepo.FindAll()
	if err != nil {
		return err
	}

	if len(*posts) == 0 {
		err := errors.New("posts do not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "all posts with ID were found successfully",
		Data:    posts,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

// ReadPost godoc
// @Summary get post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /posts/{id} [get]
func (h *BaseHandler) ReadPost(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *post == (models.Post{}) {
		err := errors.New("post with ID " + stringID + " does not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was found successfully",
		Data:    post,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

// CreatePost godoc
// @Summary create post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /posts/{id} [post]
func (h *BaseHandler) CreatePost(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	var requestPayload struct {
		UserID int    `json:"userId" xml:"userId"`
		Title  string `json:"title" xml:"title"`
		Body   string `json:"body" xml:"body"`
	}

	if err = c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *post != (models.Post{}) {
		err := errors.New("post with ID " + stringID + " exists")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	newPost := models.Post{
		ID:     ID,
		UserID: requestPayload.UserID,
		Title:  requestPayload.Title,
		Body:   requestPayload.Body,
	}

	err = h.postRepo.Create(&newPost)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was created successfully",
		Data:    newPost,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

// UpdatePost godoc
// @Summary particular update post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /posts/{id} [patch]
func (h *BaseHandler) UpdatePost(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	var requestPayload struct {
		UserID int    `json:"userId" xml:"userId"`
		Title  string `json:"title" xml:"title"`
		Body   string `json:"body" xml:"body"`
	}

	if err = c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *post == (models.Post{}) {
		err := errors.New("post with ID " + stringID + "does not exists")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
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
		return err
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was updated successfully",
		Data:    nil,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

// DeletePost godoc
// @Summary delete post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /posts/{id} [delete]
func (h *BaseHandler) DeletePost(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *post == (models.Post{}) {
		err := errors.New("post with ID " + stringID + " does not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	err = h.postRepo.Delete(post)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was deleted successfully",
	}
	return WriteResponse(c, http.StatusOK, payload)
}
