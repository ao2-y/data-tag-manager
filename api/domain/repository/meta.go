package repository

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type Meta interface {
	CreateKey(ctx context.Context, name string) (*model.MetaKey, error)
	FetchByName(ctx context.Context, name string) (*model.MetaKey, error)
}
