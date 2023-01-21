package main

import (
	"github.com/aerosystems/nix-trainee-4/pkg/mysql"

	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func main() {
	connGORM := mysql.NewClient()
	app := Config{
		DB: connGORM,
	}

	_ = app

}
