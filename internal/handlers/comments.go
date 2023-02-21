package handlers

import (
	"errors"
	"fmt"
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
		return err
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
// @Failure 404 {object} Response
// @Router /comments/{id} [post]
func (h *BaseHandler) CreateComment(c echo.Context) error {
	requestPayload := new(models.Comment)

	if err := c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	comment, err := h.commentRepo.FindByID(requestPayload.ID)
	if err != nil {
		return err
	}

	if *comment != (models.Comment{}) {
		err := fmt.Errorf("comment with ID %d exists", requestPayload.ID)
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	newComment := models.Comment{
		ID:     requestPayload.ID,
		PostId: requestPayload.PostId,
		Name:   requestPayload.Name,
		Email:  requestPayload.Email,
		Body:   requestPayload.Body,
	}

	err = h.commentRepo.Create(&newComment)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: fmt.Sprintf("comment with ID %d was created successfully", requestPayload.ID),
		Data:    newComment,
	}

	return WriteResponse(c, http.StatusCreated, payload)
}

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
		return err
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + "does not exists")
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
		return err
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was updated successfully",
		Data:    nil,
	}
	return WriteResponse(c, http.StatusOK, payload)
}

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
		return err
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exist")
		return WriteResponse(c, http.StatusNotFound, NewErrorPayload(err))
	}

	err = h.commentRepo.Delete(comment)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was deleted successfully",
	}
	return WriteResponse(c, http.StatusOK, payload)
}
