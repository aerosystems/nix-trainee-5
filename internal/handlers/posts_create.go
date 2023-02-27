package handlers

import (
	"fmt"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
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
// @Router /posts [post]
func (h *BaseHandler) CreatePost(c echo.Context) error {
	requestPayload := new(models.Post)

	if err := c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	newPost := models.Post{
		UserID: requestPayload.UserID,
		Title:  requestPayload.Title,
		Body:   requestPayload.Body,
	}

	if requestPayload.ID != 0 {
		post, err := h.postRepo.FindByID(requestPayload.ID)
		if err != nil {
			return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
		}

		if *post != (models.Post{}) {
			err := fmt.Errorf("post with ID %d exists", requestPayload.ID)
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		}

		newPost.ID = requestPayload.ID
	}

	err := h.postRepo.Create(&newPost)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: fmt.Sprintf("post with ID %d was created successfully", newPost.ID),
		Data:    newPost,
	}
	return WriteResponse(c, http.StatusCreated, payload)
}
