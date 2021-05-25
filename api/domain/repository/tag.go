package repository

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type Tag interface {
	Create(ctx context.Context, name string, parentID *uint) (*model.Tag, error)
	Remove(ctx context.Context, ID uint) (*model.Tag, error)
	FetchAll(ctx context.Context) ([]*model.Tag, error)
}
