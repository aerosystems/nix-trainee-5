package models

import "time"

type Code struct {
	ID         int       `json:"id"  xml:"id"  gorm:"<-"`
	Code       int       `json:"code"  xml:"code"  gorm:"<-"`
	UserID     int       `json:"user_id" xml:"user_id"  gorm:"<-"`
	Created    time.Time `json:"created" xml:"created"  gorm:"<-"`
	Expiration time.Time `json:"expiration" xml:"expiration"  gorm:"<-"`
	Action     string    `json:"action" xml:"action"  gorm:"<-"`
	Data       string    `json:"data" xml:"data"  gorm:"<-"`
	Used       bool      `json:"used" xml:"used"  gorm:"<-"`
}

type CodeRepository interface {
	FindAll() (*[]Code, error)
	FindByID(ID int) (*Code, error)
	Create(code *Code) error
	Update(code *Code) error
	Delete(code *Code) error
	GetByCode(Code int) (*Code, error)
	GetLastActiveCode(UserID int, Action string) (*Code, error)
	ExtendExpiration(code *Code) error
	NewCode(UserID int, Action string, Data string) (*Code, error)
}
