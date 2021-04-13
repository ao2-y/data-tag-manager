package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"errors"
)

type Meta interface {
	CreateKey(ctx context.Context, name string) (*model.MetaKey, error)
	UpdateKey(ctx context.Context, ID uint, name string) (*model.MetaKey, error)
	RemoveKey(ctx context.Context, ID uint) (*model.MetaKey, error)
	FetchKeyByID(ctx context.Context, ID uint) (*model.MetaKey, error)
}

type metaUseCase struct {
	repository repository.Meta
}

func (m *metaUseCase) FetchKeyByID(ctx context.Context, ID uint) (*model.MetaKey, error) {
	ret, err := m.repository.FetchByID(ctx, ID)
	if err != nil {
		var repError *repository.OperationError
		if errors.As(err, &repError) {
			if repError.Code == repository.ErrNotFound {
				return nil, NewResourceNorFoundError("Meta")
			}
		}
		return nil, NewInternalServerError("MetaRepository.FetchKeyByID return unknown error.", err)
	}
	return ret, nil
}

func (m *metaUseCase) UpdateKey(ctx context.Context, ID uint, name string) (*model.MetaKey, error) {
	// ハンドリングのためにユニーク制約はありつつも、チェックはやる
	ret, err := m.repository.FetchByName(ctx, name)
	if err != nil {
		var repError *repository.OperationError
		if errors.As(err, &repError) {
			switch repError.Code {
			case repository.ErrNotFound:
				// 存在しないのが正しい
				break
			default:
				return nil, NewInternalServerError("MetaRepository.FetchByName return unknown error.", err)
			}
		} else {
			return nil, NewInternalServerError("MetaRepository.FetchByName return unknown error.", err)
		}
	}
	if ret.ID != ID {
		return nil, NewValidationError(ValidationTypeExist, "Name", name, nil)
	}
	return m.repository.UpdateKey(ctx, ID, name)
}

func (m *metaUseCase) CreateKey(ctx context.Context, name string) (*model.MetaKey, error) {
	// ハンドリングのためにユニーク制約はありつつも、チェックはやる
	ret, err := m.repository.FetchByName(ctx, name)
	if err != nil {
		var repError *repository.OperationError
		if errors.As(err, &repError) {
			switch repError.Code {
			case repository.ErrNotFound:
				// 存在しないのが正しい
				break
			default:
				return nil, NewInternalServerError("MetaRepository.FetchByName return unknown error.", err)
			}
		} else {
			return nil, NewInternalServerError("MetaRepository.FetchByName return unknown error.", err)
		}
	}
	if ret != nil {
		return nil, NewValidationError(ValidationTypeExist, "Name", name, nil)
	}
	return m.repository.CreateKey(ctx, name)
}

func (m *metaUseCase) RemoveKey(ctx context.Context, ID uint) (*model.MetaKey, error) {
	// ハンドリングのために存在するか一応チェック
	_, err := m.repository.FetchByID(ctx, ID)
	if err != nil {
		var repoError *repository.OperationError
		if errors.As(err, &repoError) {
			if repoError.Code == repository.ErrNotFound {
				return nil, NewResourceNorFoundError("Meta")
			}
			return nil, NewInternalServerError("MetaRepository.FetchByID return unknown error.", err)
		}
	}
	return m.repository.RemoveKey(ctx, ID)
}

func NewMetaUseCase(repository repository.Meta) Meta {
	return &metaUseCase{
		repository: repository,
	}
}
