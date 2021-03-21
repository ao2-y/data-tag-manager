package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type Item interface {
	GetItems(ctx context.Context, ids []*uint) ([]*model.Item, error)
}

type item struct {
}

func NewItemUseCase() Item {
	return &item{}
}

func (i *item) GetItems(ctx context.Context, ids []*uint) ([]*model.Item, error) {
	var dummy []*model.Item
	desc := "fuga"
	dummy = append(dummy, &model.Item{
		ID:          1,
		Name:        "hoge",
		Description: &desc,
		Tags:        nil,
		Metas:       nil,
	})
	return dummy, nil
}
