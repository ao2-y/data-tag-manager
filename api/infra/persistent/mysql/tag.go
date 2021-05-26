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
	panic("implement me")
}

func (t *tagRepository) FetchByNameWithParentID(ctx context.Context, name string, parentID uint) ([]*model.Tag, error) {
	panic("implement me")
}

func (t *tagRepository) Remove(ctx context.Context, ID uint) (*model.Tag, error) {
	panic("implement me")
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
