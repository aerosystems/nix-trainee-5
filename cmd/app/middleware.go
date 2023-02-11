package main

import (
	"net/http"
	"os"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/handlers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *Config) AddMiddleware(e *echo.Echo) {
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}
	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))
}

func (app *Config) AuthorizationMiddleware() echo.MiddlewareFunc {
	AuthorizationConfig := echojwt.Config{
		SigningKey:     []byte(os.Getenv("ACCESS_SECRET")),
		ParseTokenFunc: app.ParseToken,
		ErrorHandler: func(c echo.Context, err error) error {
			return handlers.WriteResponse(c, http.StatusUnauthorized, handlers.NewErrorPayload(err))
		},
	}

	return echojwt.WithConfig(AuthorizationConfig)
}

func (app *Config) ParseToken(c echo.Context, auth string) (interface{}, error) {
	tokenClaims, err := app.TokensRepo.DecodeAccessToken(auth)
	if err != nil {
		return nil, err
	}

	_, err = app.TokensRepo.GetCacheValue(tokenClaims.AccessUUID)
	if err != nil {
		return nil, err
	}

	return tokenClaims, nil
}
