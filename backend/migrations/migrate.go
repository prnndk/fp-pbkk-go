package migrations

import (
	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Type{},
		&entity.Event{},
		&entity.UserTicket{},
		&entity.Pembayaran{},
	); err != nil {
		return err
	}

	return nil
}
