package main

import (
	"fmt"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/handlers"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/storage"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/mysql/mygorm"
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

	e := app.NewRouter()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", webPort)))

}
