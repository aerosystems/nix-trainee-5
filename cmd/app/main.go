package main

import (
	"fmt"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/handlers"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/storage"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/myredis"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/mysql/mygorm"
)

const webPort = 8080

type Config struct {
	BaseHandler *handlers.BaseHandler
}

// @title NIX Trainee 5-6-7-8 tasks
// @version 1.0
// @description Simple REST API for CRUD operations with Comments & Posts enities.

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	clientGORM := mygorm.NewClient()
	clientREDIS := myredis.NewClient()
	commentRepo := storage.NewCommentRepo(clientGORM)
	postRepo := storage.NewPostRepo(clientGORM)
	userRepo := storage.NewUserRepo(clientGORM, clientREDIS)
	codeRepo := storage.NewCodeRepo(clientGORM)
	tokensRepo := storage.NewTokensRepo(clientREDIS)

	app := Config{
		BaseHandler: handlers.NewBaseHandler(commentRepo, postRepo, userRepo, codeRepo, tokensRepo),
	}

	e := app.NewRouter()
	app.AddMiddleware(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", webPort)))
}
