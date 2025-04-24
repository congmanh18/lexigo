package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ChinaEntry struct {
	Base
	Traditional *string `gorm:"index:idx_traditional"`
	Simplified  *string `gorm:"index:idx_simplified"`
	Pinyin      *string
	Definition  *string
}

func (e *ChinaEntry) TableName() string {
	return "china_entries"
}
