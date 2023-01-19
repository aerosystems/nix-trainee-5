package post

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	gorm *gorm.DB
}

func (r *Repository) Create(post Post) {
	result := r.gorm.Create(&post)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
