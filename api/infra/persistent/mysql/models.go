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
		ID   uint `gorm:"primaryKey"`
		Name string
		//CreatedAt   time.Time
		//UpdatedAt   time.Time
		//DeletedAt   gorm.DeletedAt `gorm:"index"`
	}

	Tags struct {
		ID          uint `gorm:"primaryKey"`
		Name        string
		ParentTagID uint
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
		MetaKeys []*ItemTemplateMetaKeys `gorm:"foreignKey:ItemTemplateID"`
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

func itemTemplatesToDomain(templates []*ItemTemplates) []*model.ItemTemplate {
	ret := make([]*model.ItemTemplate, len(templates), len(templates))
	for i, v := range templates {
		ret[i] = itemTemplateToDomain(v)
	}
	return ret
}

func itemTemplateToDomain(template *ItemTemplates) *model.ItemTemplate {
	return &model.ItemTemplate{
		ID:       template.ID,
		Name:     template.Name,
		MetaKeys: itemTemplateMetaKeysToDomain(template.MetaKeys),
	}
}

func itemTemplateMetaKeysToDomain(keys []*ItemTemplateMetaKeys) []*model.MetaKey {
	ret := make([]*model.MetaKey, len(keys), len(keys))
	for i, v := range keys {
		ret[i] = itemTemplateMetaKeyToDomain(v)
	}
	return ret
}

func itemTemplateMetaKeyToDomain(key *ItemTemplateMetaKeys) *model.MetaKey {
	return &model.MetaKey{
		ID:   key.MetaKeys.ID,
		Name: key.MetaKeys.Name,
	}
}

func metaKeyToDomain(key *MetaKeys) *model.MetaKey {
	return &model.MetaKey{
		ID:   key.ID,
		Name: key.Name,
	}
}

func metaKeysToDomain(keys []*MetaKeys) []*model.MetaKey {
	ret := make([]*model.MetaKey, len(keys), len(keys))
	for i, v := range keys {
		ret[i] = metaKeyToDomain(v)
	}
	return ret
}

func tagToDomain(tag *Tags) *model.Tag {
	return &model.Tag{
		ID:          tag.ID,
		Name:        tag.Name,
		ParentTagID: tag.ParentTagID,
	}
}

func tagsToDomain(tags []*Tags) []*model.Tag {
	ret := make([]*model.Tag, len(tags), len(tags))
	for i, v := range tags {
		ret[i] = tagToDomain(v)
	}
	return ret
}
