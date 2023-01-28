package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *BaseHandler) ReadComments(c echo.Context) error {
	comments, err := h.commentRepo.FindAll()
	if err != nil {
		return err
	}

	if len(*comments) == 0 {
		err := errors.New("comments do not exist")
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	payload := Response{
		Error:   false,
		Message: "all comments with ID were found successfully",
		Data:    comments,
	}
	return MakeResponse(c, http.StatusOK, payload)
}

func (h *BaseHandler) ReadComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exist")
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was found successfully",
		Data:    comment,
	}
	return MakeResponse(c, http.StatusOK, payload)
}

func (h *BaseHandler) CreateComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	requestPayload := new(models.Comment)

	if err = c.Bind(&requestPayload); err != nil {
		return MakeResponse(c, http.StatusBadRequest, NewError(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *comment != (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " exists")
		return MakeResponse(c, http.StatusBadRequest, NewError(err))
	}

	newComment := models.Comment{
		Id:     ID,
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
		Message: "comment with ID " + stringID + " was created successfully",
		Data:    newComment,
	}

	return MakeResponse(c, http.StatusOK, payload)
}

func (h *BaseHandler) UpdateComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	requestPayload := new(models.Comment)

	if err = c.Bind(&requestPayload); err != nil {
		return MakeResponse(c, http.StatusBadRequest, NewError(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + "does not exists")
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	newComment := comment
	newComment.Id = ID

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
	return MakeResponse(c, http.StatusOK, payload)
}

func (h *BaseHandler) DeleteComment(c echo.Context) error {
	stringID := c.Param("id")

	ID, err := strconv.Atoi(stringID)
	if err != nil {
		return MakeResponse(c, http.StatusBadRequest, NewError(err))
	}

	comment, err := h.commentRepo.FindByID(ID)
	if err != nil {
		return err
	}

	if *comment == (models.Comment{}) {
		err := errors.New("comment with ID " + stringID + " does not exist")
		return MakeResponse(c, http.StatusNotFound, NewError(err))
	}

	err = h.commentRepo.Delete(comment)
	if err != nil {
		return err
	}

	payload := Response{
		Error:   false,
		Message: "comment with ID " + stringID + " was deleted successfully",
	}
	return MakeResponse(c, http.StatusOK, payload)
}
