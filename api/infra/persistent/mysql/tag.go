package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"ao2-y/data-tag-manager/infra/persistent/inmemory"
	"context"
	"gorm.io/gorm"
)

type tagRepository struct {
	db    *gorm.DB
	cache inmemory.Cache
}

func (t *tagRepository) Create(ctx context.Context, name string, parentID *uint) (*model.Tag, error) {
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
