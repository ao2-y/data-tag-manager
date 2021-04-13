package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddMetaKey(ctx context.Context, input *model.AddMetaKeyInput) (*model.AddMetaKeyPayload, error) {
	useCaseRet, err := r.MetaUseCase.CreateKey(ctx, input.Name)
	if err != nil {
		return nil, createError("AddMetaKey operation failed", err)
	}
	return &model.AddMetaKeyPayload{
		ClientMutationID: input.ClientMutationID,
		MetaKey: &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(useCaseRet.ID),
			Name: useCaseRet.Name,
		},
	}, nil
}

func (r *mutationResolver) UpdateMetaKey(ctx context.Context, input *model.UpdateMetaKeyInput) (*model.UpdateMetaKeyPayload, error) {
	innerID, err := model.KeyMeta.ToInternalID(input.ID)
	if err != nil {
		return nil, createError("UpdateMetaKey ID validation failed.", nil)
	}
	useCaseRet, err := r.MetaUseCase.UpdateKey(ctx, innerID, input.Name)
	if err != nil {
		return nil, createError("UpdateMetaKey operation failed", err)
	}
	return &model.UpdateMetaKeyPayload{
		ClientMutationID: input.ClientMutationID,
		MetaKey: &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(useCaseRet.ID),
			Name: useCaseRet.Name,
		},
	}, nil
}

func (r *mutationResolver) RemoveMetaKey(ctx context.Context, input *model.RemoveMetaKeyInput) (*model.RemoveMetaKeyPayload, error) {
	innerID, err := model.KeyMeta.ToInternalID(input.ID)
	if err != nil {
		return nil, createError("RemoveMetaKey ID validation failed.", nil)
	}
	useCaseRet, err := r.MetaUseCase.RemoveKey(ctx, innerID)
	if err != nil {
		return nil, createError("UpdateMetaKey operation failed", err)
	}
	return &model.RemoveMetaKeyPayload{
		ClientMutationID: input.ClientMutationID,
		MetaKey: &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(useCaseRet.ID),
			Name: useCaseRet.Name,
		},
	}, nil
}

func (r *mutationResolver) AddMetaToItem(ctx context.Context, input *model.AddMetaToItemInput) (*model.AddMetaToItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveMetaToItem(ctx context.Context, input *model.RemoveMetaToItemInput) (*model.RemoveMetaToItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MetaKeys(ctx context.Context) ([]*model.MetaKey, error) {
	panic(fmt.Errorf("not implemented"))
}
