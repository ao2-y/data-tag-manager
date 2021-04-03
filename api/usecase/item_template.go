package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"context"
)

type ItemTemplateUsecase interface {
	GetAll(ctx context.Context) []*model.ItemTemplate
}
