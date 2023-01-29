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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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
