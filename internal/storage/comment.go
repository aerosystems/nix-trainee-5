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

func (r *commentRepo) FindAll() (*[]models.Comment, error) {
	var comments []models.Comment
	r.db.Find(&comments)
	return &comments, nil
}

func (r *commentRepo) FindByID(ID int) (*models.Comment, error) {
	var comment models.Comment
	r.db.Find(&comment, ID)
	return &comment, nil
}

func (r *commentRepo) Create(comment *models.Comment) error {
	result := r.db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *commentRepo) Update(comment *models.Comment) error {
	result := r.db.Save(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *commentRepo) Delete(comment *models.Comment) error {
	result := r.db.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
