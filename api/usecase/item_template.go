package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"fmt"
)

type ItemTemplate interface {
	FetchAll(ctx context.Context) ([]*model.ItemTemplate, error)
	FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error)
	Create(ctx context.Context, Name string, MetaKeyIDs []*uint) (*model.ItemTemplate, error)
	Update(ctx context.Context, ID uint, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error)
	Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error)
}

func NewItemTemplateUseCase(
	itemTemplateRepository repository.ItemTemplate,
	metaRepository repository.Meta,
) ItemTemplate {
	return &itemTemplate{
		itemTemplateRepository: itemTemplateRepository,
		metaRepository:         metaRepository,
	}
}

type itemTemplate struct {
	itemTemplateRepository repository.ItemTemplate
	metaRepository         repository.Meta
}

func (i *itemTemplate) FetchAll(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	return nil, fmt.Errorf("not implement")
}

func (i *itemTemplate) Create(ctx context.Context, Name string, MetaKeyIDs []*uint) (*model.ItemTemplate, error) {
	// MetaKeyIDの生存確認
	template, err := i.itemTemplateRepository.Create(ctx, Name, MetaKeyIDs)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (i *itemTemplate) Update(ctx context.Context, ID uint, Name string, MetaKeyIDs []string) (*model.ItemTemplate, error) {
	panic("implement me")
}

func (i *itemTemplate) Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error) {
	panic("implement me")
}