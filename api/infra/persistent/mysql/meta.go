package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
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
			return nil, repository.NewStoreOperationError(repository.StoreOperationCodeNotFound, nil)
		}
		return nil, repository.NewStoreOperationError(repository.StoreOperationCodeUnkownError, result.Error)
	}
	return metaKeyToDomain(metaKeys), nil
}

func (m *metaRepository) CreateKey(ctx context.Context, name string) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		Name: name,
	}
	result := m.db.WithContext(ctx).Create(metaKey)
	if result.Error != nil {
		result.Rollback()
		return nil, repository.NewStoreOperationError(repository.StoreOperationCodeUnkownError, result.Error)
	}
	result.Commit()
	return metaKeyToDomain(metaKey), nil
}

func NewMetaRepository(db *gorm.DB) repository.Meta {
	return &metaRepository{
		db: db,
	}
}
