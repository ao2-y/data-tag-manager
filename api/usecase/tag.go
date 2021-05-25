package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
)

type Tag interface {
	Create(ctx context.Context, name string, parentID uint) (*model.Tag, error)
	Remove(ctx context.Context, ID uint) (*model.Tag, error)
	GetAll(ctx context.Context) ([]*model.Tag, error)
	GetByID(ctx context.Context, ID uint) (*model.Tag, error)
	GetByIDWithParent(ctx context.Context, ID uint) (*model.Tag, error)
}

type tagUseCase struct {
	repository repository.Tag
}

func (t *tagUseCase) Create(ctx context.Context, name string, parentID uint) (*model.Tag, error) {
	panic("implement me")
}

func (t *tagUseCase) Remove(ctx context.Context, ID uint) (*model.Tag, error) {
	// 子が存在するIDは削除させない
	// Itemに紐づいているIDは削除させない
	panic("implement me")
}

func (t *tagUseCase) GetAll(ctx context.Context) ([]*model.Tag, error) {
	panic("implement me")
}

func (t *tagUseCase) GetByID(ctx context.Context, ID uint) (*model.Tag, error) {
	panic("implement me")
}

func NewTagUseCase(repository repository.Tag) Tag {
	return &tagUseCase{repository: repository}
}
