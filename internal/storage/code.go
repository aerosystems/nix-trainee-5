package storage

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"gorm.io/gorm"
)

type codeRepo struct {
	db *gorm.DB
}

func NewCodeRepo(db *gorm.DB) *codeRepo {
	return &codeRepo{
		db: db,
	}
}

func (r *codeRepo) FindAll() (*[]models.Code, error) {
	var codes []models.Code
	r.db.Find(&codes)
	return &codes, nil
}

func (r *codeRepo) FindByID(ID int) (*models.Code, error) {
	var code models.Code
	r.db.Find(&code, ID)
	return &code, nil
}

func (r *codeRepo) Create(code *models.Code) error {
	result := r.db.Create(&code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *codeRepo) Update(code *models.Code) error {
	result := r.db.Save(&code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *codeRepo) Delete(code *models.Code) error {
	result := r.db.Delete(&code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *codeRepo) GetByCode(Code int) (*models.Code, error) {
	var code models.Code
	r.db.Find(&code, Code)
	return &code, nil
}

func (r *codeRepo) GetLastActiveCode(UserID int, Action string) (*models.Code, error) {
	var code models.Code
	r.db.Find(&code, UserID, Action)
	return &code, nil
}

func (r *codeRepo) ExtendExpiration(code *models.Code) error {
	codeExpMinutes, err := strconv.Atoi(os.Getenv("CODE_EXP_MINUTES"))
	if err != nil {
		return err
	}
	code.Expiration = time.Now().Add(time.Minute * time.Duration(codeExpMinutes))
	result := r.db.Save(&code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateCode generation new code
func (r *codeRepo) NewCode(UserID int, Action string, Data string) (*models.Code, error) {
	codeExpMinutes, err := strconv.Atoi(os.Getenv("CODE_EXP_MINUTES"))
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	XXXXXX := rand.Intn(max-min+1) + min
	code := models.Code{
		Code:       XXXXXX,
		UserID:     UserID,
		Created:    time.Now(),
		Expiration: time.Now().Add(time.Minute * time.Duration(codeExpMinutes)),
		Action:     Action,
		Data:       Data,
		Used:       false,
	}

	result := r.db.Create(&code)
	if result.Error != nil {
		return nil, result.Error
	}
	return &code, nil
}
