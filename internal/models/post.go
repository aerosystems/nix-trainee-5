package models

type Post struct {
	ID     int    `json:"id" gorm:"<-"`
	UserID int    `json:"userId" gorm:"<-"`
	Title  string `json:"title" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type PostRepository interface {
	FindAll() (*[]Post, error)
	FindByID(ID int) (*Post, error)
	Create(post *Post) error
	Update(post *Post) error
	Delete(post *Post) error
}
