package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aerosystems/nix-trainee-4/internal/handlers"
	"github.com/aerosystems/nix-trainee-4/internal/storage"
	"github.com/aerosystems/nix-trainee-4/pkg/mysql/mygorm"
)

const webPort = 8080

type Config struct {
	BaseHandler *handlers.BaseHandler
}

func main() {
	clientGORM := mygorm.NewClient()
	commentRepo := storage.NewCommentRepo(clientGORM)
	postRepo := storage.NewPostRepo(clientGORM)

	app := Config{
		BaseHandler: handlers.NewBaseHandler(commentRepo, postRepo),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	log.Printf("Starting service on port %d\n", webPort)
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
