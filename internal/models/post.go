package models

import (
	"log"

	"gorm.io/gorm"
)

type Post struct {
	Id     int    `json:"id" gorm:"<-"`
	UserId int    `json:"userId" gorm:"<-"`
	Title  string `json:"title" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type PostRepository struct {
	gorm *gorm.DB
}

func (pr *PostRepository) Create(post Post) {
	result := pr.gorm.Create(&post)
	if result.Error != nil {
		log.Println(result.Error)
	}
}
