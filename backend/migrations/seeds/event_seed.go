package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/prnndk/final-project-golang-pbkk/entity"
	"gorm.io/gorm"
)

func ListEventSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/events.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listEvent []entity.Event
	if err := json.Unmarshal(jsonData, &listEvent); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Event{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Event{}); err != nil {
			return err
		}
	}

	for _, data := range listEvent {
		var event entity.Event
		var event_type entity.Type
		err := db.Where(&entity.Event{Name: data.Name}).First(&event).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		data_entity := db.Where(&entity.Type{Name: data.TypeID}).First(&event_type).Error
		if data_entity != nil && !errors.Is(data_entity, gorm.ErrRecordNotFound) {
			return err
		}

		data.Date = event.Date
		data.TypeID = event_type.ID.String()

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
