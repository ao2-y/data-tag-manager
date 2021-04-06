package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type itemTemplate struct {
	db *gorm.DB
}

func (i *itemTemplate) Create(ctx context.Context, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error) {
	metaKeys := make([]ItemTemplateMetaKeys, len(metaKeyIDs))
	for i, v := range metaKeyIDs {
		metaKeys[i] = ItemTemplateMetaKeys{
			MetaKeyID: *v,
		}
	}
	it := &ItemTemplates{
		Name:     name,
		MetaKeys: metaKeys,
	}
	result := i.db.WithContext(ctx).Create(it)
	if result.Error != nil {
		result.Rollback()
		return nil, fmt.Errorf("item template create error. %w", result.Error)
	}
	result.Commit()

	return i.FetchByID(ctx, it.ID)
}

func (i *itemTemplate) FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	itemTemplates := ItemTemplates{
		ID: ID,
	}
	result := i.db.WithContext(ctx).First(&itemTemplates)
	if result.Error != nil {
		return nil, fmt.Errorf("item template FetchByID error.:%w", result.Error)
	}

	ret := itemTemplateToDomain(itemTemplates)
	return &ret, nil
}

func NewItemTemplateRepository(db *gorm.DB) repository.ItemTemplate {
	return &itemTemplate{
		db: db,
	}
}
