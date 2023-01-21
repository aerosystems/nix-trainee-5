package models

import (
	"log"

	"gorm.io/gorm"
)

type Comment struct {
	Id     int    `json:"id" gorm:"<-"`
	PostId int    `json:"postId" gorm:"<-"`
	Name   string `json:"name" gorm:"<-"`
	Email  string `json:"email" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type CommentRepository struct {
	gorm *gorm.DB
}

func (cr *CommentRepository) Create(comment Comment) {
	result := cr.gorm.Create(&comment)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
