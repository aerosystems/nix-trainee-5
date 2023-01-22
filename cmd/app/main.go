package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aerosystems/nix-trainee-4/internal/handlers"
	"github.com/aerosystems/nix-trainee-4/internal/storage"
	"github.com/aerosystems/nix-trainee-4/pkg/mysql/mygorm"
)

func main() {
	clientGORM := mygorm.NewClient()
	commentRepo := storage.NewCommentRepo(clientGORM)
	_ = handlers.NewBaseHandler(commentRepo)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
