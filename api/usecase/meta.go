package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"fmt"
)

type Meta interface {
	CreateKey(ctx context.Context, name string) (*model.MetaKey, error)
}

type metaUseCase struct {
	repository repository.Meta
}

func (m *metaUseCase) CreateKey(ctx context.Context, name string) (*model.MetaKey, error) {
	// TODO チェックではなくDBの制約にする？
	ret, err := m.repository.FetchByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("CreateKey:FetchByName Failed:%w", err)
	}
	if ret != nil {
		return nil, fmt.Errorf("dupulicated name")
	}
	return m.repository.CreateKey(ctx, name)
}

func NewMetaUseCase(repository repository.Meta) Meta {
	return &metaUseCase{
		repository: repository,
	}
}
