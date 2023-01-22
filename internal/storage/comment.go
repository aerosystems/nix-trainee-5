package storage

import (
	"github.com/aerosystems/nix-trainee-4/internal/models"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) FindByID(ID int) (*models.Comment, error) {
	var comment models.Comment
	r.db.Find(&comment, ID)
	return &comment, nil
}

func (r *commentRepo) Save(comment *models.Comment) error {
	result := r.db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
