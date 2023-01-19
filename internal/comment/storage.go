package comment

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	gorm *gorm.DB
}

func (r *Repository) Create(comment Comment) {
	result := r.gorm.Create(&comment)
	if result.Error != nil {
		log.Println(result.Error)
	}
}

func NewRepository(gorm *gorm.DB) *Repository {
	return &Repository{
		gorm: gorm,
	}
}
