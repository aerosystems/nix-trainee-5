package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

// UpdateComment godoc
// @Summary particular update comment by ID
// @Tags comments
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param	id	path	int	true "Comment ID"
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response{data=models.Comment}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 404 {object} Response
// @Router /comments/{id} [patch]
func (h *BaseHandler) UpdateComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	requestPayload := new(models.Comment)

	if err = c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exists")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	newComment := comment
	newComment.ID = ID

	if requestPayload.PostId != 0 {
		newComment.PostId = requestPayload.PostId
	}
	if requestPayload.Name != "" {
		newComment.Name = requestPayload.Name
	}
	if requestPayload.Email != "" {
		newComment.Email = requestPayload.Email
	}
	if requestPayload.Body != "" {
		newComment.Body = requestPayload.Body
	}

	err = h.commentRepo.Update(newComment)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was updated successfully",
		Data:    nil,
	}
	return WriteResponse(c, http.StatusOK, payload)
}
