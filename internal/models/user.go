package models

import "time"

type User struct {
	ID       int       `json:"id" xml:"id"  gorm:"<-"`
	Email    string    `json:"email" xml:"email"  gorm:"<-"`
	Password string    `json:"-" xml:"_"  gorm:"<-"`
	Role     string    `json:"role" xml:"role"  gorm:"<-"`
	Created  time.Time `json:"created" xml:"created"  gorm:"<-"`
	Updated  time.Time `json:"updated" xml:"updated"  gorm:"<-"`
	Active   bool      `json:"active" xml:"active"  gorm:"<-"`
}

type UserRepository interface {
	FindAll() (*[]User, error)
	FindByID(ID int) (*User, error)
	FindByEmail(Email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	ResetPassword(user *User, password string) error
	PasswordMatches(user *User, plainText string) (bool, error)
}
