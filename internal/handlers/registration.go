package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/helpers"
	"github.com/labstack/echo/v4"
)

func (h *BaseHandler) Registration(c echo.Context) error {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&requestPayload); err != nil {
		return WriteResponse(c, http.StatusBadRequest, err)
	}

	addr, err := helpers.ValidateEmail(requestPayload.Email)
	if err != nil {
		err = errors.New("email is not valid")
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	email := helpers.NormalizeEmail(addr)

	// Minimum of one small case letter
	// Minimum of one upper case letter
	// Minimum of one digit
	// Minimum of one special character
	// Minimum 8 characters length
	err = helpers.ValidatePassword(requestPayload.Password)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	var payload Response

	//checking if email is existing
	user, err := h.userRepo.FindByEmail(email)
	fmt.Println(email, user, "\n", err)
	if user != nil {
		fmt.Println("dfsdfsdfsdf")
		if user.IsActive {
			err = errors.New("email already exists")
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		} else {
			// updating password for inactive user
			err := h.userRepo.ResetPassword(user, requestPayload.Password)
			if err != nil {
				return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
			}

			code, _ := h.codeRepo.GetLastIsActiveCode(user.ID, "registration")

			if code == nil {
				// generating confirmation code
				_, err = h.codeRepo.NewCode(user.ID, "registration", "")
				if err != nil {
					return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
				}
			} else {
				// extend expiration code and return previous active code
				h.codeRepo.ExtendExpiration(code)
				_ = code.Code
			}

			payload := Response{
				Error:   false,
				Message: fmt.Sprintf("Updated user with Id: %d", user.ID),
				Data:    nil,
			}

			return WriteResponse(c, http.StatusAccepted, payload)
		}
	}

	// creating new inactive user
	newUser := models.User{
		Email:    email,
		Password: requestPayload.Password,
		Role:     "user",
	}
	err = h.userRepo.Create(&newUser)

	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	// generating confirmation code
	code, err := h.codeRepo.NewCode(newUser.ID, "registration", "")
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	payload = Response{
		Error:   false,
		Message: fmt.Sprintf("Registered user with Id: %d. Confirmation code: %d", newUser.ID, code.Code),
		Data:    nil,
	}

	return WriteResponse(c, http.StatusOK, payload)
}
