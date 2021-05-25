package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddTag(ctx context.Context, input *model.AddTagInput) (*model.AddTagPaylod, error) {
	var parentID uint
	var err error
	if input.ParentID != nil {
		parentID, err = model.KeyTag.ToInternalID(*input.ParentID)
		if err != nil {
			return nil, newGraphqlError("AddTag ParentID validation failed.", err)
		}
	}
	useCaseRet, err := r.TagUseCase.Create(ctx, input.Name, parentID)
	if err != nil {
		return nil, newGraphqlError("AddTag operation failed", err)
	}

	return &model.AddTagPaylod{
		ClientMutationID: input.ClientMutationID,
		Tag: &model.Tag{
			ID:     model.KeyTag.ToExternalID(useCaseRet.ID),
			Parent: nil,
			Name:   useCaseRet.Name,
		},
	}, nil
}

func (r *mutationResolver) RemoveTag(ctx context.Context, input *model.RemoveTagInput) (*model.RemoveTagPayload, error) {
	innerID, err := model.KeyTag.ToInternalID(input.ID)
	if err != nil {
		return nil, newGraphqlError("RemoveTag ID validation failed", nil)
	}
	useCaseRet, err := r.TagUseCase.Remove(ctx, innerID)
	if err != nil {
		return nil, newGraphqlError("RemoveTag operation failed", err)
	}
	return &model.RemoveTagPayload{
		ClientMutationID: input.ClientMutaionID,
		Tag: &model.Tag{
			ID:     model.KeyTag.ToExternalID(useCaseRet.ID),
			Parent: nil,
			Name:   useCaseRet.Name,
		},
	}, nil
}

func (r *mutationResolver) AddTagToItem(ctx context.Context, input *model.AddTagToItemInput) (*model.AddTagToItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveTagToItem(ctx context.Context, input *model.RemoveTagToItemInput) (*model.RemoveTagToItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}
