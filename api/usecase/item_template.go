package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
	"fmt"
)

type ItemTemplate interface {
	FetchAll(ctx context.Context) ([]*model.ItemTemplate, error)
	Fetch(ctx context.Context, ID uint) (*model.ItemTemplate, error)
	Create(ctx context.Context, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error)
	Update(ctx context.Context, ID uint, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error)
	Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error)
}

func NewItemTemplateUseCase() ItemTemplate {
	return &itemTemplate{}
}

type itemTemplate struct{}

func (i *itemTemplate) FetchAll(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) Fetch(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	return nil, fmt.Errorf("not implement")
}

func (i *itemTemplate) Create(ctx context.Context, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error) {
	return nil, fmt.Errorf("not implement")
}

func (i *itemTemplate) Update(ctx context.Context, ID uint, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	panic("implement me")
}
