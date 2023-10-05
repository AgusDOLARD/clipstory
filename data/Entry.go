package data

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	database "github.com/AgusDOLARD/clipstory/db"
)

type Entry struct {
	gorm.Model
	Value string `gorm:"value"`
}

func (e *Entry) Save() error {
	e.Value = strings.TrimSpace(e.Value)
	if e.Value == ""{
		return errors.New("Entry is an empty string")
	}
	err := database.Db.Create(e).Error
	if err != nil {
		return fmt.Errorf("Entry Save Error: %w", err)
	}
	return nil
}

func GetEntries() ([]Entry, error) {
	var entries []Entry
	err := database.Db.Order("created_at DESC").Find(&entries).Error
	if err != nil {
		return nil, fmt.Errorf("Entry GetEntries Error: %w", err)
	}
	return entries, nil
}
