package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

// ReadComments godoc
// @Summary get all comments
// @Tags comments
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 200 {object} Response{data=[]models.Comment}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 404 {object} Response
// @Router /comments [get]
func (h *BaseHandler) ReadComments(c echo.Context) error {
	comments, err := h.commentRepo.FindAll()
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	if len(*comments) == 0 {
		err := errors.New("comments do not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "all comments with ID were found successfully",
		Data:    comments,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

// ReadComment godoc
// @Summary get comment by ID
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
// @Router /comments/{id} [get]
func (h *BaseHandler) ReadComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was found successfully",
		Data:    comment,
	}
	return WriteResponse(c, http.StatusOK, payload)
}
