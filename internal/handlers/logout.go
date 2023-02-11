package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *BaseHandler) Logout(c echo.Context) error {
	payload := Response{
		Error:   true,
		Message: "sdfsdf",
		Data:    nil,
	}
	return WriteResponse(c, http.StatusAccepted, payload)
	// accessTokenClaims, ok := r.Context().Value(contextKey("accessTokenClaims")).(*AccessTokenClaims)
	// if !ok {
	// 	_ = app.errorJSON(w, errors.New("token is untracked"), http.StatusUnauthorized)
	// 	return
	// }

	// err := app.dropCacheTokens(*accessTokenClaims)
	// if err != nil {
	// 	_ = app.errorJSON(w, err, http.StatusUnauthorized)
	// 	return
	// }

	// payload := jsonResponse{
	// 	Error:   false,
	// 	Message: fmt.Sprintf("User %s successfully logged out", accessTokenClaims.AccessUUID),
	// 	Data:    accessTokenClaims,
	// }

	// _ = app.writeJSON(w, http.StatusAccepted, payload)
}
