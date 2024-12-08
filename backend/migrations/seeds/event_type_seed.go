package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

func ListEventTypeSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/event_type.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listType []entity.Type
	if err := json.Unmarshal(jsonData, &listType); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Type{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Type{}); err != nil {
			return err
		}
	}

	for _, data := range listType {
		var event_type entity.Type
		err := db.Where(&entity.Type{Name: data.Name}).First(&event_type).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
