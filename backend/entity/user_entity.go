package entity

import (
	"github.com/google/uuid"
	"github.com/prnndk/final-project-golang-pbkk/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
