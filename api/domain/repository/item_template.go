package repository

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type ItemTemplate interface {
	Create(ctx context.Context, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error)
	Update(ctx context.Context, ID uint, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error)
	Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error)
	FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error)
	FetchAll(ctx context.Context) ([]*model.ItemTemplate, error)
}
