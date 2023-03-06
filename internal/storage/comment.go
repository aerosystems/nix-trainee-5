package storage

import (
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) FindAll() (*[]models.Comment, error) {
	var comments []models.Comment
	r.db.Find(&comments)
	return &comments, nil
}

func (r *CommentRepo) FindByID(ID int) (*models.Comment, error) {
	var comment models.Comment
	r.db.Find(&comment, ID)
	return &comment, nil
}

func (r *CommentRepo) Create(comment *models.Comment) error {
	result := r.db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CommentRepo) Update(comment *models.Comment) error {
	result := r.db.Save(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CommentRepo) Delete(comment *models.Comment) error {
	result := r.db.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
