package repository

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type ItemTemplate interface {
	//FetchByID(ctx context.Context) ([]*model.ItemTemplate,error)
	Create(ctx context.Context, name string, metaKeyIDs []*uint) (*model.ItemTemplate, error)
}
