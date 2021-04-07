package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type metaRepository struct {
	db *gorm.DB
}

func (m *metaRepository) FetchByName(ctx context.Context, name string) (*model.MetaKey, error) {
	metaKeys := &MetaKeys{}
	result := m.db.WithContext(ctx).Where("name = ?", name).First(metaKeys)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("MetaKey fetch failed:%w", result.Error)
	}
	return metaKeyToDomain(metaKeys), nil
}

func (m *metaRepository) CreateKey(ctx context.Context, name string) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		Name: name,
	}
	result := m.db.Create(metaKey)
	if result.Error != nil {
		result.Rollback()
		return nil, fmt.Errorf("CreateKey error:%w", result.Error)
	}
	result.Commit()
	return metaKeyToDomain(metaKey), nil
}

func NewMetaRepository(db *gorm.DB) repository.Meta {
	return &metaRepository{
		db: db,
	}
}
