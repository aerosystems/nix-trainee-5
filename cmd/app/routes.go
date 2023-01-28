package main

import (
	"github.com/labstack/echo/v4"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/v1/comments", app.BaseHandler.ReadComments)
	e.GET("/v1/comments/:id", app.BaseHandler.ReadComment)
	e.POST("/v1/comments/:id", app.BaseHandler.CreateComment)
	e.PATCH("/v1/comments/:id", app.BaseHandler.UpdateComment)
	e.DELETE("/v1/comments/:id", app.BaseHandler.DeleteComment)
	e.GET("/v1/posts", app.BaseHandler.ReadPosts)
	e.GET("/v1/posts/:id", app.BaseHandler.ReadPost)
	e.POST("/v1/posts/:id", app.BaseHandler.CreatePost)
	e.PATCH("/v1/posts/:id", app.BaseHandler.UpdatePost)
	e.DELETE("/v1/posts/:id", app.BaseHandler.DeletePost)

	return e
}
