package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"ao2-y/data-tag-manager/infra/persistent/inmemory"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type metaRepository struct {
	db    *gorm.DB
	cache inmemory.Cache
}

func (m *metaRepository) FetchByIDs(ctx context.Context, IDs []*uint) ([]*model.MetaKey, error) {
	// TODO とりあえずキャッシュ見てない
	var metas []*MetaKeys
	err := m.db.WithContext(ctx).Where("ID IN ?", IDs).Find(&metas).Error
	if err != nil {
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	return metaKeysToDomain(metas), nil
}

func (m *metaRepository) FetchByID(ctx context.Context, iD uint) (*model.MetaKey, error) {
	if ret := m.cache.Restore(m.cacheKeyID(iD)); ret != nil {
		if meta, ok := ret.(*model.MetaKey); ok {
			return meta, nil
		}
	}

	metaKeys := &MetaKeys{ID: iD}
	err := m.db.WithContext(ctx).First(metaKeys).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	modelMeta := metaKeyToDomain(metaKeys)
	m.cache.Store(m.cacheKeyID(iD), modelMeta)
	return modelMeta, nil
}

func (m *metaRepository) FetchByName(ctx context.Context, name string) (*model.MetaKey, error) {
	if ret := m.cache.Restore(m.cacheKeyName(name)); ret != nil {
		if meta, ok := ret.(*model.MetaKey); ok {
			return meta, nil
		}
	}

	metaKeys := &MetaKeys{}
	result := m.db.WithContext(ctx).Where("name = ?", name).First(metaKeys)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, result.Error)
	}
	modelMeta := metaKeyToDomain(metaKeys)
	m.cache.Store(m.cacheKeyName(name), modelMeta)
	return modelMeta, nil
}

func (m *metaRepository) CreateKey(ctx context.Context, name string) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		Name: name,
	}
	tx := m.db.WithContext(ctx).Begin()
	err := tx.WithContext(ctx).Create(metaKey).Error
	if err != nil {
		tx.Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	tx.Commit()
	modelMeta := metaKeyToDomain(metaKey)
	m.cache.Store(m.cacheKeyID(metaKey.ID), modelMeta)
	m.cache.Store(m.cacheKeyName(metaKey.Name), modelMeta)
	return modelMeta, nil
}

func (m *metaRepository) UpdateKey(ctx context.Context, iD uint, name string) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		ID: iD,
	}
	tx := m.db.WithContext(ctx).Begin()
	err := tx.WithContext(ctx).Set("gorm:query_option", "FOR UPDATE").First(metaKey).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	// 古いキャッシュを削除
	m.cache.Delete(m.cacheKeyID(iD))
	m.cache.Delete(m.cacheKeyName(metaKey.Name))

	metaKey.Name = name

	err = tx.WithContext(ctx).Save(metaKey).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	tx.WithContext(ctx).Commit()

	modelMeta := metaKeyToDomain(metaKey)
	m.cache.Store(m.cacheKeyID(iD), modelMeta)
	m.cache.Store(m.cacheKeyName(name), modelMeta)
	return modelMeta, nil
}

func (m *metaRepository) RemoveKey(ctx context.Context, ID uint) (*model.MetaKey, error) {
	metaKey := &MetaKeys{
		ID: ID,
	}
	tx := m.db.WithContext(ctx).Begin()
	err := tx.WithContext(ctx).Set("gorm:query_option", "FOR UPDATE").First(metaKey).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	// Cache削除
	m.cache.Delete(m.cacheKeyID(metaKey.ID))
	m.cache.Delete(m.cacheKeyName(metaKey.Name))

	err = tx.WithContext(ctx).Delete(metaKey).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	tx.WithContext(ctx).Commit()
	return metaKeyToDomain(metaKey), nil
}

func (m *metaRepository) cacheKeyID(iD uint) string {
	return fmt.Sprintf("MetaKeyID:%v", iD)
}

func (m *metaRepository) cacheKeyName(name string) string {
	return fmt.Sprintf("MetaKeyName:%s", name)
}

func NewMetaRepository(db *gorm.DB, isLocalCacheEnabled bool) repository.Meta {
	var cache inmemory.Cache
	if isLocalCacheEnabled {
		cache = inmemory.NewInMemoryCache()
	} else {
		cache = inmemory.NewNoCache()
	}
	return &metaRepository{
		db:    db,
		cache: cache,
	}
}
