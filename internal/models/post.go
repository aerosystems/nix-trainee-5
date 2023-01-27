package models

type Post struct {
	ID     int    `json:"id" xml:"id" gorm:"<-"`
	UserID int    `json:"userId" xml:"userId" gorm:"<-"`
	Title  string `json:"title" xml:"title" gorm:"<-"`
	Body   string `json:"body" xml:"body" gorm:"<-"`
}

type PostRepository interface {
	FindAll() (*[]Post, error)
	FindByID(ID int) (*Post, error)
	Create(post *Post) error
	Update(post *Post) error
	Delete(post *Post) error
}
