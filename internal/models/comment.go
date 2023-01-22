package models

type Comment struct {
	Id     int    `json:"id" gorm:"<-"`
	PostId int    `json:"postId" gorm:"<-"`
	Name   string `json:"name" gorm:"<-"`
	Email  string `json:"email" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type CommentRepository interface {
	FindByID(ID int) (*Comment, error)
	Save(comment *Comment) error
}