package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

// DeleteComment godoc
// @Summary delete comment by ID
// @Tags comments
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Comment ID"
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 404 {object} Response
// @Router /comments/{id} [delete]
func (h *BaseHandler) DeleteComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	err = h.commentRepo.Delete(comment)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was deleted successfully",
	}
	return WriteResponse(c, http.StatusOK, payload)
}
