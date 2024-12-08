package migrations

import (
	"github.com/prnndk/final-project-golang-pbkk/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	// if err := seeds.ListUserSeeder(db); err != nil {
	// 	return err
	// }

	if err := seeds.ListEventTypeSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListEventSeeder(db); err != nil {
		return err
	}

	return nil
}
