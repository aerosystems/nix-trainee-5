package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Public routes
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/v1/comments", app.BaseHandler.ReadComments)
	mux.Get("/v1/comments/{id}", app.BaseHandler.ReadComment)
	mux.Post("/v1/comments/{id}", app.BaseHandler.CreateComment)
	mux.Patch("/v1/comments/{id}", app.BaseHandler.UpdateComment)
	mux.Delete("/v1/comments/{id}", app.BaseHandler.DeleteComment)
	mux.Get("/v1/posts", app.BaseHandler.ReadPosts)
	mux.Get("/v1/posts/{id}", app.BaseHandler.ReadPost)
	mux.Post("/v1/posts/{id}", app.BaseHandler.CreatePost)
	mux.Patch("/v1/posts/{id}", app.BaseHandler.UpdatePost)
	mux.Delete("/v1/posts/{id}", app.BaseHandler.DeletePost)

	return mux
}
