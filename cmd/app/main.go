package main

import (
	"fmt"
	"os"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/handlers"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/storage"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/myredis"
	"github.com/aerosystems/nix-trainee-5-6-7-8/pkg/mysql/mygorm"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const webPort = 8080

type Config struct {
	BaseHandler       *handlers.BaseHandler
	GoogleOauthConfig *oauth2.Config
	TokensRepo        models.TokensRepository
}

// @title NIX Trainee 5-6-7-8 tasks
// @version 1.0
// @description Simple REST API for CRUD operations with Comments & Posts entities.

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

	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/v1/callback/google",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	app := Config{
		BaseHandler: handlers.NewBaseHandler(googleOauthConfig,
			commentRepo,
			postRepo,
			userRepo,
			codeRepo,
			tokensRepo,
		),
		TokensRepo: tokensRepo,
	}

	e := app.NewRouter()
	app.AddMiddleware(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", webPort)))
}
