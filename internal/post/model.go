package post

type Post struct {
	Id     int    `json:"id" gorm:"<-"`
	UserId int    `json:"userId" gorm:"<-"`
	Title  string `json:"title" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}
