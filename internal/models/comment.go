package models

type Comment struct {
	Id     int    `json:"id" gorm:"<-"`
	PostId int    `json:"postId" gorm:"<-"`
	Name   string `json:"name" gorm:"<-"`
	Email  string `json:"email" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type CommentRepository interface {
	FindAll() (*[]Comment, error)
	FindByID(ID int) (*Comment, error)
	Create(comment *Comment) error
	Update(comment *Comment) error
	Delete(comment *Comment) error
}
