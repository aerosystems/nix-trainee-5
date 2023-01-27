package models

type Comment struct {
	Id     int    `json:"id" xml:"id" gorm:"<-"`
	PostId int    `json:"postId" xml:"postId" gorm:"<-"`
	Name   string `json:"name" xml:"name" gorm:"<-"`
	Email  string `json:"email" xml:"email" gorm:"<-"`
	Body   string `json:"body" xml:"body" gorm:"<-"`
}

type CommentRepository interface {
	FindAll() (*[]Comment, error)
	FindByID(ID int) (*Comment, error)
	Create(comment *Comment) error
	Update(comment *Comment) error
	Delete(comment *Comment) error
}
