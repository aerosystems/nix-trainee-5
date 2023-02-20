package models

type Post struct {
	ID     int    `json:"id" xml:"id" gorm:"<-;primaryKey" example:"61"`
	UserID int    `json:"userId" xml:"userId" gorm:"<-" example:"7"`
	Title  string `json:"title" xml:"title" gorm:"<-" example:"voluptatem doloribus..."`
	Body   string `json:"body" xml:"body" gorm:"<-" example:"dolore maxime saepe..."`
}

type PostRepository interface {
	FindAll() (*[]Post, error)
	FindByID(ID int) (*Post, error)
	Create(post *Post) error
	Update(post *Post) error
	Delete(post *Post) error
}
