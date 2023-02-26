package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
)

// ReadPosts godoc
// @Summary get all posts
// @Tags posts
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response{data=[]models.Post}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
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
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response{data=models.Post}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
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
