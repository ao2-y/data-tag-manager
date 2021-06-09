package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"errors"
	"fmt"
)

type ItemTemplate interface {
	FetchAll(ctx context.Context) ([]*model.ItemTemplate, error)
	FetchByID(ctx context.Context, ID uint) (*model.ItemTemplate, error)
	Create(ctx context.Context, name string, MetaKeyIDs []*uint) (*model.ItemTemplate, error)
	UpdateName(ctx context.Context, ID uint, name string) (*model.ItemTemplate, error)
	UpdateMetaKeys(ctx context.Context, ID uint, MetaKeyIDs []*uint) (*model.ItemTemplate, error)
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

func (i *itemTemplate) Create(ctx context.Context, name string, MetaKeyIDs []*uint) (*model.ItemTemplate, error) {
	metas, err := i.metaRepository.FetchByIDs(ctx, MetaKeyIDs)
	if len(metas) == len(MetaKeyIDs) {
		// 件数が一致してないなら取得できてないやつが存在しているのでエラー
		return nil, NewValidationError(ValidationTypeExist, "MetaKeyIDs", MetaKeyIDs, nil)
	}
	template, err := i.itemTemplateRepository.Create(ctx, name, MetaKeyIDs)
	if err != nil {
		return nil, NewInternalServerError("ItemTemplate create operation failed", err)
	}
	return template, nil
}

func (i *itemTemplate) UpdateName(ctx context.Context, ID uint, name string) (*model.ItemTemplate, error) {
	// TODO
	// IDの存在チェック→変更なしならそのまま返す
	// nameの重複チェック
	panic("implement me")
}

func (i *itemTemplate) UpdateMetaKeys(ctx context.Context, ID uint, MetaKeyIDs []*uint) (*model.ItemTemplate, error) {
	// TODO
	// IDの存在チェック→変更なしならそのまま返す
	// MetaKeyIDsの存在チェック
	panic("implement me")
}

func (i *itemTemplate) Remove(ctx context.Context, ID uint) (*model.ItemTemplate, error) {

	it, err := i.itemTemplateRepository.FetchByID(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("ItemTemplate")
			}
		}
		return nil, NewInternalServerError("ItemTemplate Remove usecase failed.", err)
	}
	_, err = i.itemTemplateRepository.Remove(ctx, ID)
	if err != nil {
		return nil, NewInternalServerError("ItemTemplate usecase Remove operation failed.", err)
	}
	return it, nil
}
