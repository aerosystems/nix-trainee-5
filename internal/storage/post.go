package storage

import (
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) FindAll() (*[]models.Post, error) {
	var posts []models.Post
	r.db.Find(&posts)
	return &posts, nil
}

func (r *PostRepo) FindByID(ID int) (*models.Post, error) {
	var post models.Post
	r.db.Find(&post, ID)
	return &post, nil
}

func (r *PostRepo) Create(post *models.Post) error {
	result := r.db.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostRepo) Update(post *models.Post) error {
	result := r.db.Save(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostRepo) Delete(post *models.Post) error {
	result := r.db.Delete(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
