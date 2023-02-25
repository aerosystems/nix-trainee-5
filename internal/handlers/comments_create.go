package handlers

import (
	"fmt"
	"net/http"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

// CreateComment godoc
// @Summary create comment by ID
// @Tags comments
// @Accept  json
// @Accept  xml
// @Produce application/json
// @Produce application/xml
// @Param comment body models.Comment true "raw request body"
// @Param Authorization header string true "should contain Access Token, with the Bearer started"
// @Success 201 {object} Response{data=models.Comment}
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Router /comments [post]
func (h *BaseHandler) CreateComment(c echo.Context) error {
	requestPayload := new(models.Comment)

	if err := c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	newComment := models.Comment{
		PostId: requestPayload.PostId,
		Name:   requestPayload.Name,
		Email:  requestPayload.Email,
		Body:   requestPayload.Body,
	}

	if requestPayload.ID != 0 {
		comment, err := h.commentRepo.FindByID(requestPayload.ID)
		if err != nil {
			return err
		}

		if *comment != (models.Comment{}) {
			err := fmt.Errorf("comment with ID %d exists", requestPayload.ID)
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		}

		newComment.ID = requestPayload.ID
	}

	err := h.commentRepo.Create(&newComment)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: fmt.Sprintf("comment with ID %d was created successfully", newComment.ID),
		Data:    newComment,
	}

	return WriteResponse(c, http.StatusCreated, payload)
}
