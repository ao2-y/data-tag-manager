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

type itemTemplate struct {
	db    *gorm.DB
	cache inmemory.Cache
}

func (i *itemTemplate) Update(ctx context.Context, ID uint, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) FetchAll(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) Create(ctx context.Context, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error) {
	metaKeys := make([]*ItemTemplateMetaKeys, len(metaKeyIDs))
	for i, v := range metaKeyIDs {
		metaKeys[i] = &ItemTemplateMetaKeys{
			MetaKeyID: *v,
		}
	}
	it := &ItemTemplates{
		Name:     name,
		MetaKeys: metaKeys,
	}
	err := i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Create(it).Error
		if err != nil {
			return repository.NewOperationError(repository.ErrUnknown, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	ret := itemTemplateToDomain(it)
	i.cache.Store(i.cacheKeyID(ret.ID), ret)
	i.cache.Store(i.cacheKeyName(ret.Name), ret)
	return ret, nil
}

func (i *itemTemplate) FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	if ret := i.cache.Restore(i.cacheKeyID(ID)); ret != nil {
		if itemTemplate, ok := ret.(*model.ItemTemplate); ok {
			return itemTemplate, nil
		}
	}
	itemTemplates := &ItemTemplates{
		ID: ID,
	}
	err := i.db.WithContext(ctx).First(itemTemplates).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.NewOperationError(repository.ErrNotFound, nil)
		}
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	ret := itemTemplateToDomain(itemTemplates)
	return ret, nil
}

func (i *itemTemplate) cacheKeyID(iD uint) string {
	return fmt.Sprintf("ItemTemplateID:%v", iD)
}

func (i *itemTemplate) cacheKeyName(name string) string {
	return fmt.Sprintf("ItemTemplateName:%s", name)
}

func NewItemTemplateRepository(db *gorm.DB, isLocalCacheEnabled bool) repository.ItemTemplate {
	var cache inmemory.Cache
	if isLocalCacheEnabled {
		cache = inmemory.NewInMemoryCache()
	} else {
		cache = inmemory.NewNoCache()
	}
	return &itemTemplate{
		db:    db,
		cache: cache,
	}
}
