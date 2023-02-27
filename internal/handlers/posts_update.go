package handlers

import (
	"errors"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// UpdatePost godoc
// @Summary particular update post by ID
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Post ID"
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 404 {object} Response
// @Router /posts/{id} [patch]
func (h *BaseHandler) UpdatePost(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	requestPayload := new(models.Post)

	if err = c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	post, err := h.postRepo.FindByID(ID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	if *post == (models.Post{}) {
		err := errors.New("post with ID " + stringID + " does not exists")
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
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "post with ID " + stringID + " was updated successfully",
		Data:    nil,
	}
	return WriteResponse(c, http.StatusOK, payload)
}
