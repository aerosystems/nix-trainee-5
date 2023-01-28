package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

type BaseHandler struct {
	commentRepo models.CommentRepository
	postRepo    models.PostRepository
}

// Response is the type used for sending JSON around
type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewBaseHandler(commentRepo models.CommentRepository, postRepo models.PostRepository) *BaseHandler {
	return &BaseHandler{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func NewError(err error) Response {
	return Response{
		Error:   true,
		Message: err.Error(),
	}
}

func MakeResponse(c echo.Context, statusCode int, payload any) error {
	acceptHeaders := c.Request().Header["Accept"]
	if Contains(acceptHeaders, "application/xml") {
		return c.XML(statusCode, payload)
	}
	return c.JSON(statusCode, payload)
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// writeResponse takes a response status code and arbitrary data and writes a xml/json response to the client in depends of Header Accept
func writeResponse(w http.ResponseWriter, r *http.Request, status int, data any, headers ...http.Header) error {
	var out []byte
	var err error
	switch r.Header.Get("Accept") {
	case "application/xml":
		out, err = xml.MarshalIndent(data, "", " ")
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/xml")
	default:
		out, err = json.Marshal(data)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a xml/json response to the client in depends of Header Accept
func writeError(w http.ResponseWriter, r *http.Request, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload Response
	payload.Error = true
	payload.Message = err.Error()

	return writeResponse(w, r, statusCode, payload)
}

// readRequest tries to read the body of a request and converts it into JSON/XML in depends of Header Accept
func readRequest(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	switch r.Header.Get("Content-Type") {
	case "application/xml":
		body, _ := io.ReadAll(r.Body)
		err := xml.Unmarshal(body, data)
		if err != nil {
			return err
		}
	case "application/json":
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(data)
		if err != nil {
			return err
		}
		err = dec.Decode(&struct{}{})
		if err != io.EOF {
			return errors.New("body must have only a single json value")
		}
	default:
		return errors.New("unknown Content-Type. Supported types are application/xml and application/json")
	}
	return nil
}
