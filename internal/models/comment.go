package models

type Comment struct {
	ID     int    `json:"id" xml:"id" gorm:"<-;primaryKey" example:"302"`
	PostId int    `json:"postId" xml:"postId" gorm:"<-" example:"61"`
	Name   string `json:"name" xml:"name" gorm:"<-" example:"quia voluptatem..."`
	Email  string `json:"email" xml:"email" gorm:"<-" example:"lindsey@caitlyn.net"`
	Body   string `json:"body" xml:"body" gorm:"<-" example:"fuga aut est delectus..."`
}

type CommentRepository interface {
	FindAll() (*[]Comment, error)
	FindByID(ID int) (*Comment, error)
	Create(comment *Comment) error
	Update(comment *Comment) error
	Delete(comment *Comment) error
}
