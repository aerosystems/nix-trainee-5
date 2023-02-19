package models

type Comment struct {
	Id     int    `json:"id" xml:"id" gorm:"column:id;primaryKey" example:"302"`
	PostId int    `json:"postId" xml:"postId" gorm:"column:post_id" example:"61"`
	Name   string `json:"name" xml:"name" gorm:"column:name" example:"quia voluptatem..."`
	Email  string `json:"email" xml:"email" gorm:"column:email" example:"lindsey@caitlyn.net"`
	Body   string `json:"body" xml:"body" gorm:"column:body" example:"fuga aut est delectus..."`
}

type CommentRepository interface {
	FindAll() (*[]Comment, error)
	FindByID(ID int) (*Comment, error)
	Create(comment *Comment) error
	Update(comment *Comment) error
	Delete(comment *Comment) error
}
