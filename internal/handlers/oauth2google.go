package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func (h *BaseHandler) Oauth2GoogleLogin(c echo.Context) error {
	b := make([]byte, 16)
	rand.Read(b)
	oauthStateString := base64.URLEncoding.EncodeToString(b)

	cookie := new(http.Cookie)

	cookie.Name = "oauthstate"
	cookie.Value = oauthStateString
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false	
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)

	url := h.googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusMovedPermanently, url)
}

func (h *BaseHandler) Oauth2GoogleCallback(c echo.Context) error {
	oauthstateCookie, err := c.Cookie("oauthstate")
	if err != nil {
		customErr := fmt.Errorf("something wrong with web browser... %s", err.Error())
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(customErr))
	}

	if c.FormValue("state") != oauthstateCookie.Value {
		err := errors.New("hmm... seems like CSRF attack")
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))

	}

	content, err := h.getUserInfo(c.FormValue("code"))
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	var googleUserInfo UserInfo
	if err := json.Unmarshal(content, &googleUserInfo); err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	user, err := h.userRepo.FindByGoogleID(googleUserInfo.Sub)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		}
	}
	message := "successfull Authorization with Google OAuth2.0"
	// case #1: user not found in storage
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// creating new active user
		newUser := models.User{
			GoogleID: googleUserInfo.Sub,
			Email:    googleUserInfo.Email,
			Role:     "user",
			IsActive: true,
		}
		err = h.userRepo.Create(&newUser)
		if err != nil {
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		}

		user, err = h.userRepo.FindByGoogleID(googleUserInfo.Sub)
		if err != nil {
			return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
		}
		message = "first successfull Authorization with Google OAuth2.0"
	}
	// case #2: if user found in storage
	// create pair of JWT tokens
	ts, err := h.tokensRepo.CreateToken(user.ID)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	// add refresh token UUID to cache
	err = h.tokensRepo.CreateCacheKey(user.ID, ts)
	if err != nil {
		return WriteResponse(c, http.StatusBadRequest, NewErrorPayload(err))
	}

	tokens := TokensResponseBody{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	payload := Response{
		Error:   false,
		Message: message,
		Data:    tokens,
	}

	return WriteResponse(c, http.StatusAccepted, payload)
}

func (h *BaseHandler) getUserInfo(code string) ([]byte, error) {
	token, err := h.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
