package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *BaseHandler) Logout(c echo.Context) error {
	accessTokenClaims, ok := c.Get("user").(*models.AccessTokenClaims)
	if !ok {
		err := errors.New("token is untracked")
		return WriteResponse(c, http.StatusUnauthorized, NewErrorPayload(err))
	}

	err := h.tokensRepo.DropCacheTokens(*accessTokenClaims)
	if err != nil {
		return WriteResponse(c, http.StatusUnauthorized, NewErrorPayload(err))
	}

	payload := Response{
		Error:   false,
		Message: fmt.Sprintf("User %s successfully logged out", accessTokenClaims.AccessUUID),
		Data:    accessTokenClaims,
	}
	return WriteResponse(c, http.StatusAccepted, payload)
}
