package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	"context"
	"errors"
	"gopkg.in/go-playground/colors.v1"
)

type Tag interface {
	Create(ctx context.Context, name string, color string, parentID uint) (*model.Tag, error)
	Update(ctx context.Context, ID uint, name, color string) (*model.Tag, error)
	Remove(ctx context.Context, ID uint) (*model.Tag, error)
	GetAll(ctx context.Context) ([]*model.Tag, error)
	GetByID(ctx context.Context, ID uint) (*model.Tag, error)
	GetByIDWithParent(ctx context.Context, ID uint) (*model.TagWithParent, error)
}

type tagUseCase struct {
	repository repository.Tag
}

func (t *tagUseCase) GetByIDWithParent(ctx context.Context, ID uint) (*model.TagWithParent, error) {
	tag, err := t.repository.FetchByID(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			}
		}
		return nil, NewInternalServerError("TagUseCase GetByIDWithParent failed", err)
	}
	var parentTag *model.Tag
	if tag.ParentTagID > 0 {
		p, err := t.repository.FetchByID(ctx, tag.ParentTagID)
		if err != nil {
			return nil, NewInternalServerError("TagUseCase Parent Fetch failed", err)
		}
		parentTag = p
	}
	return &model.TagWithParent{
		Tag: model.Tag{
			ID:          tag.ID,
			Name:        tag.Name,
			ParentTagID: tag.ParentTagID,
		},
		Parent: parentTag,
	}, nil
}

func (t *tagUseCase) Create(ctx context.Context, name string, color string, parentID uint) (*model.Tag, error) {

	// カラーコードの形式チェック
	_, err := colors.ParseHEX(color)
	if err != nil {
		return nil, NewValidationError(ValidationTypeColorCode, "Color", color, err)
	}
	if parentID > 0 {
		// Parentの生存確認
		parent, err := t.repository.FetchByID(ctx, parentID)
		if err != nil {
			var opeError *repository.OperationError
			if errors.As(err, opeError) {
				switch opeError.Code {
				case repository.ErrNotFound:
					return nil, NewValidationError(ValidationTypeExist, "ParentID", parentID, nil)
				default:
					return nil, NewInternalServerError("Tag create usecase. get parent failed", err)
				}
			}
			return nil, NewInternalServerError("Tag create usecase. get parent failed", err)
		}
		if parent.ParentTagID > 0 {
			// ParentにしたいTagが子。孫は作れない。
			return nil, NewValidationError(ValidationTypeIsChild, "ParentID", parentID, nil)
		}
	}

	sameNameTags, err := t.repository.FetchByNameWithParentID(ctx, name, parentID)
	if err != nil {
		// FIXME errors.Asでやる
		return nil, NewInternalServerError("Tag create usecase error", err)
	}
	if len(sameNameTags) > 0 {
		// 同名のタグが存在するためエラーで返す
		return nil, NewValidationError(ValidationTypeDuplicated, "name", name, nil)
	}

	tag, err := t.repository.Create(ctx, name, color, parentID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			return nil, NewInternalServerError("Tag create usecase. create failed.", err)
		}
		return nil, NewInternalServerError("Tag create failed", err)
	}
	return tag, nil
}

func (t tagUseCase) Update(ctx context.Context, ID uint, name, color string) (*model.Tag, error) {
	_, err := t.repository.FetchByID(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			}
		}
		return nil, NewInternalServerError("Tag update usecase. FetchByID failed.", err)
	}
	tag, err := t.repository.Update(ctx, ID, name, color)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			}
		}
		return nil, NewInternalServerError("Tag update usecase. Update failed.", err)
	}
	return tag, nil
}

func (t *tagUseCase) Remove(ctx context.Context, ID uint) (*model.Tag, error) {
	_, err := t.repository.FetchByID(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			default:
				return nil, NewInternalServerError("Tag remove usecase. FetchByID failed.", err)
			}
		}
		return nil, NewInternalServerError("Tag remove usecase. FetchByID failed.", err)
	}
	// 子が存在するIDは削除させない
	children, err := t.repository.FetchByParentID(ctx, ID)
	if err != nil {
		return nil, NewInternalServerError("Tag remove usecase. FetchByParendID failed.", err)
	}
	if len(children) > 0 {
		return nil, NewValidationError(ValidationTypeUsed, "ID", ID, nil)
	}

	// TODO FIXME Itemに紐づいているIDは削除させない

	tag, err := t.repository.Remove(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			default:
				return nil, NewInternalServerError("Tag remove usecase. Remove failed.", err)
			}
		}
		return nil, NewInternalServerError("Tag remove usecase. Remove failed.", err)
	}

	return tag, nil
}

func (t *tagUseCase) GetAll(ctx context.Context) ([]*model.Tag, error) {
	tags, err := t.repository.FetchAll(ctx)
	if err != nil {
		return nil, NewInternalServerError("TagUseCase GetAll operation failed.", err)
	}
	return tags, nil
}

func (t *tagUseCase) GetByID(ctx context.Context, ID uint) (*model.Tag, error) {
	tag, err := t.repository.FetchByID(ctx, ID)
	if err != nil {
		var opeError *repository.OperationError
		if errors.As(err, opeError) {
			switch opeError.Code {
			case repository.ErrNotFound:
				return nil, NewResourceNorFoundError("Tag")
			}
		}
		return nil, NewInternalServerError("Tag GetByID usecase failed.", err)
	}
	return tag, nil
}

func NewTagUseCase(repository repository.Tag) Tag {
	return &tagUseCase{repository: repository}
}
