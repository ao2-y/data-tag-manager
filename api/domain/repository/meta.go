package repository

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type Meta interface {
	CreateKey(ctx context.Context, name string) (*model.MetaKey, error)
	UpdateKey(ctx context.Context, ID uint, name string) (*model.MetaKey, error)
	RemoveKey(ctx context.Context, ID uint) (*model.MetaKey, error)
	FetchByName(ctx context.Context, name string) (*model.MetaKey, error)
	FetchByID(ctx context.Context, ID uint) (*model.MetaKey, error)
	//FetchByIDs(ctx context.Context, IDs ...uint) ([]*model.MetaKey, error)
}
