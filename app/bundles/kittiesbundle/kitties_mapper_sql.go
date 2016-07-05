package kittiesbundle

import (
	"github.com/jinzhu/gorm"
)

// KittiesMapperSQL define a SQL mapper
type KittiesMapperSQL struct {
	db *gorm.DB
}

// NewKittiesSQLMapper instance
func NewKittiesSQLMapper(db *gorm.DB) *KittiesMapperSQL {
	return &KittiesMapperSQL{
		db: db,
	}
}

// FindAll kitties in database
func (m *KittiesMapperSQL) FindAll() ([]Kitty, error) {
	var kitties []Kitty

	m.db.Find(&kitties)

	return kitties, nil
}
