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

func (m *metaRepository) FetchByID(ctx context.Context, ID uint) (*model.MetaKey, error) {
	metaKeys := &MetaKeys{ID: ID}
	err := m.db.WithContext(ctx).First(metaKeys).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	return metaKeyToDomain(metaKeys), nil
}

func (m *metaRepository) FetchByName(ctx context.Context, name string) (*model.MetaKey, error) {
	metaKeys := &MetaKeys{}
	result := m.db.WithContext(ctx).Where("name = ?", name).First(metaKeys)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, result.Error)
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
		return nil, repository.NewOperationError(repository.ErrUnknown, result.Error)
	}
	result.Commit()
	return metaKeyToDomain(metaKey), nil
}

func (m *metaRepository) UpdateKey(ctx context.Context, ID uint, name string) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		ID:   ID,
		Name: name,
	}
	result := m.db.WithContext(ctx).Save(metaKey)
	if result.Error != nil {
		result.Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, result.Error)
	}
	return metaKeyToDomain(metaKey), nil
}

func (m *metaRepository) RemoveKey(ctx context.Context, ID uint) (*model.MetaKey, error) {

	metaKey := &MetaKeys{
		ID: ID,
	}
	tx := m.db.WithContext(ctx).Begin()
	err := tx.Set("gorm:query_option", "FOR UPDATE").First(metaKey).Error
	if err != nil {
		tx.Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	err = tx.Delete(metaKey).Error
	if err != nil {
		tx.Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	tx.Commit()
	return metaKeyToDomain(metaKey), nil
}

func NewMetaRepository(db *gorm.DB) repository.Meta {
	return &metaRepository{
		db: db,
	}
}
