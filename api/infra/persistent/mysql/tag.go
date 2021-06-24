package mysql

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"ao2-y/data-tag-manager/infra/persistent/inmemory"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type tagRepository struct {
	db    *gorm.DB
	cache inmemory.Cache
}

func (t *tagRepository) Update(ctx context.Context, ID uint, name string, color string, parentID uint) (*model.Tag, error) {
	tag := &Tags{
		ID:          ID,
		Name:        name,
		Color:       color,
		ParentTagID: parentID,
	}
	err := t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Save(tag).Error
		if err != nil {
			return repository.NewOperationError(repository.ErrUnknown, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	modelTag := tagToDomain(tag)
	t.cache.Store(t.cacheKeyID(modelTag.ID), modelTag)
	t.cache.Store(t.cacheKeyName(modelTag.Name), modelTag)
	// TODO FIXME ParentIDに関しては同じParentIDを持っている群をキャッシュしているので読みなおし処理必要
	// 暫定で同じparentIDを持つキャッシュを削除
	t.cache.Delete(t.cacheKeyParentID(modelTag.ParentTagID))
	return modelTag, nil
}

func (t *tagRepository) Create(ctx context.Context, name string, color string, parentID uint) (*model.Tag, error) {
	tag := &Tags{
		Name:        name,
		Color:       color,
		ParentTagID: parentID,
	}
	err := t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Create(tag).Error
		if err != nil {
			return repository.NewOperationError(repository.ErrUnknown, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	modelTag := tagToDomain(tag)
	t.cache.Store(t.cacheKeyID(modelTag.ID), modelTag)
	t.cache.Store(t.cacheKeyName(modelTag.Name), modelTag)
	// TODO FIXME ParentIDに関しては同じParentIDを持っている群をキャッシュしているので読みなおし処理必要
	// 暫定で同じparentIDを持つキャッシュを削除
	t.cache.Delete(t.cacheKeyParentID(modelTag.ParentTagID))
	return modelTag, nil
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
	err := t.db.WithContext(ctx).Where("parent_tag_id = ?", ID).Find(&tags).Error
	if err != nil {
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}

	modelTags := tagsToDomain(tags)
	t.cache.Store(t.cacheKeyParentID(ID), modelTags)
	return modelTags, nil
}

func (t *tagRepository) FetchByNameWithParentID(ctx context.Context, name string, parentID uint) ([]*model.Tag, error) {
	// キャッシュから検索しにくいのでキャッシュ見ない
	var tags []*Tags
	err := t.db.WithContext(ctx).Where("parent_tag_id = @parent and name = @name", sql.NamedArg{
		Name:  "parent",
		Value: parentID,
	}, sql.NamedArg{
		Name:  "name",
		Value: name,
	}).Find(&tags).Error
	if err != nil {
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	return tagsToDomain(tags), nil
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
	var tags []*Tags
	err := t.db.WithContext(ctx).Find(&tags).Error
	if err != nil {
		return nil, repository.NewOperationError(repository.ErrUnknown, err)
	}
	modelTags := tagsToDomain(tags)
	return modelTags, nil
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
