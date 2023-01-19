package main

import (
	"github.com/aerosystems/nix-trainee-4/internal/comment"
	"github.com/aerosystems/nix-trainee-4/pkg/client/gorm"
)

func main() {
	clientGORM := gorm.NewClient()
	repository := comment.NewRepository(clientGORM)
	_ = repository

}
