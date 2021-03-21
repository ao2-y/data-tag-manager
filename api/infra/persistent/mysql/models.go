package mysql

import (
	"database/sql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type (
	Items struct {
		ID          uint `gorm:"primaryKey"`
		Name        string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}

	MetaKeys struct {
		ID          uint `gorm:"primaryKey"`
		ParentTagId sql.NullInt64
		Level       uint8
		Name        string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
	}

	Tags struct {
		ID        uint `gorm:"primaryKey"`
		ItemID    uint
		MetaKeyID uint
		Value     string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	ItemTags struct {
		ID        uint `gorm:"primaryKey"`
		ItemID    uint
		TagID     uint
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	ItemTemplates struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	ItemTemplateMetaKeys struct {
		ID             uint `gorm:"primaryKey"`
		ItemTemplateID uint
		MetaKeyID      uint
		CreatedAt      time.Time
		UpdatedAt      time.Time
		DeletedAt      gorm.DeletedAt `gorm:"index"`
	}
)
