package handlers

import (
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

type BaseHandler struct {
	commentRepo models.CommentRepository
	postRepo    models.PostRepository
}

// Response is the type used for sending JSON around
type Response struct {
	Error   bool   `json:"error" xml:"error"`
	Message string `json:"message" xml:"message"`
	Data    any    `json:"data,omitempty" xml:"data,omitempty"`
}

func NewBaseHandler(commentRepo models.CommentRepository, postRepo models.PostRepository) *BaseHandler {
	return &BaseHandler{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

// WriteResponse takes a response status code and arbitrary data and writes a xml/json response to the client in depends of Header Accept
func WriteResponse(c echo.Context, statusCode int, payload any) error {
	acceptHeaders := c.Request().Header["Accept"]
	if Contains(acceptHeaders, "application/xml") {
		return c.XML(statusCode, payload)
	}
	return c.JSON(statusCode, payload)
}

func NewErrorPayload(err error) Response {
	return Response{
		Error:   true,
		Message: err.Error(),
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
