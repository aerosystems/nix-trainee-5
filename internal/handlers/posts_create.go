package handlers

import (
	"errors"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"

	"strconv"
)

// CreatePost godoc
// @Summary create post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 201 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
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
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
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
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was created successfully",
		Data:    newPost,
	}
	return WriteResponse(c, http.StatusCreated, payload)
}
