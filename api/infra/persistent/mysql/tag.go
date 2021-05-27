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

type tagRepository struct {
	db    *gorm.DB
	cache inmemory.Cache
}

func (t *tagRepository) Create(ctx context.Context, name string, parentID uint) (*model.Tag, error) {
	panic("implement me")
}

func (t *tagRepository) FetchByID(ctx context.Context, ID uint) (*model.Tag, error) {
	if ret := t.cache.Restore(t.cacheKeyID(ID)); ret != nil {
		if tag, ok := ret.(*model.Tag); ok {
			return tag, nil
		}
	}

	tag := &Tags{
		ID: ID,
	}
	err := t.db.WithContext(ctx).First(tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	modelTag := tagToDomain(tag)
	t.cache.Store(t.cacheKeyID(ID), modelTag)
	return modelTag, nil
}

func (t *tagRepository) FetchByParentID(ctx context.Context, ID uint) ([]*model.Tag, error) {
	if ID == 0 {
		// ID==0は親なしなので空Sliceで返す
		return []*model.Tag{}, nil
	}
	if ret := t.cache.Restore(t.cacheKeyParentID(ID)); ret != nil {
		if tags, ok := ret.([]*model.Tag); ok {
			return tags, nil
		}
	}
	var tags []*Tags
	err := t.db.WithContext(ctx).Where("parent_id = ?", ID).Find(&tags).Error
	if err != nil {
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	modelTags := tagsToDomain(tags)
	t.cache.Store(t.cacheKeyParentID(ID), modelTags)
	return modelTags, nil
}

func (t *tagRepository) FetchByNameWithParentID(ctx context.Context, name string, parentID uint) ([]*model.Tag, error) {
	panic("implement me")
}

func (t *tagRepository) Remove(ctx context.Context, ID uint) (*model.Tag, error) {
	tag := &Tags{ID: ID}
	tx := t.db.WithContext(ctx).Begin()
	err := tx.WithContext(ctx).Set("gorm:query_option", "FOR UPDATE").First(tag).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	// Cache削除
	t.cache.Delete(t.cacheKeyID(ID))
	t.cache.Delete(t.cacheKeyName(tag.Name))

	err = tx.WithContext(ctx).Delete(tag).Error
	if err != nil {
		tx.WithContext(ctx).Rollback()
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	tx.WithContext(ctx).Commit()
	return tagToDomain(tag), nil
}

func (t *tagRepository) FetchAll(ctx context.Context) ([]*model.Tag, error) {
	panic("implement me")
}

func NewTagRepository(db *gorm.DB, isLocalCacheEnabled bool) repository.Tag {
	var cache inmemory.Cache
	if isLocalCacheEnabled {
		cache = inmemory.NewInMemoryCache()
	} else {
		cache = inmemory.NewNoCache()
	}
	return &tagRepository{
		db:    db,
		cache: cache,
	}
}

func (t *tagRepository) cacheKeyID(iD uint) string {
	return fmt.Sprintf("TagID:%v", iD)
}

func (t *tagRepository) cacheKeyParentID(iD uint) string {
	return fmt.Sprintf("TagParentID:%v", iD)
}

func (t *tagRepository) cacheKeyName(name string) string {
	return fmt.Sprintf("TagName:%s", name)
}
