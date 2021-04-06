package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	_ "gorm.io/driver/mysql"
)

type (
	Items struct {
		ID          uint `gorm:"primaryKey"`
		Name        string
		Description *string
		//CreatedAt   time.Time
		//UpdatedAt   time.Time
		//DeletedAt   gorm.DeletedAt `gorm:"index"`
		ItemTags []ItemTags `gorm:"foreignKey:ItemID"`
	}

	MetaKeys struct {
		ID           uint `gorm:"primaryKey"`
		ParentMetaId *uint
		Level        uint8
		Name         string
		//CreatedAt   time.Time
		//UpdatedAt   time.Time
		//DeletedAt   gorm.DeletedAt `gorm:"index"`
		ParentMeta *MetaKeys
	}

	Tags struct {
		ID        uint `gorm:"primaryKey"`
		ItemID    uint
		MetaKeyID uint
		Value     string
		//CreatedAt time.Time
		//UpdatedAt time.Time
		//DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	ItemTags struct {
		ID     uint `gorm:"primaryKey"`
		ItemID uint
		TagID  uint
		//CreatedAt time.Time
		//UpdatedAt time.Time
		//DeletedAt gorm.DeletedAt `gorm:"index"`
		Tag Tags `gorm:"foreignKey:TagID;references:TagID"`
	}

	ItemTemplates struct {
		ID   uint `gorm:"primaryKey"`
		Name string
		//CreatedAt time.Time
		//UpdatedAt time.Time
		//DeletedAt gorm.DeletedAt `gorm:"index"`
		MetaKeys []ItemTemplateMetaKeys `gorm:"foreignKey:ItemTemplateID"`
	}

	ItemTemplateMetaKeys struct {
		ID             uint `gorm:"primaryKey"`
		ItemTemplateID uint
		MetaKeyID      uint
		//CreatedAt      time.Time
		//UpdatedAt      time.Time
		//DeletedAt      gorm.DeletedAt `gorm:"index"`
		MetaKeys MetaKeys `gorm:"foreignKey:MetaKeyID"`
	}
)

func itemToDomain(item Items) model.Item {
	return model.Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Tags:        nil,
		Metas:       nil,
	}
}

func itemTemplateToDomain(template ItemTemplates) model.ItemTemplate {
	return model.ItemTemplate{
		ID:       template.ID,
		Name:     template.Name,
		MetaKeys: metaKeysToDomain(template.MetaKeys),
	}
}

func metaKeysToDomain(keys []ItemTemplateMetaKeys) []model.MetaKey {
	ret := make([]model.MetaKey, 0, len(keys))
	for i, v := range keys {
		ret[i] = metaKeyToDomain(v)
	}
	return ret
}

func metaKeyToDomain(key ItemTemplateMetaKeys) model.MetaKey {
	return model.MetaKey{
		ID:   key.MetaKeys.ID,
		Name: key.MetaKeys.Name,
	}
}
