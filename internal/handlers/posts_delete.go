package handlers

import (
	"errors"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// DeletePost godoc
// @Summary delete post by ID
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
