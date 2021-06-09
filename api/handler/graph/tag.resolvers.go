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

	var parent *model.Tag
	if parentID > 0 {
		parentRet, err := r.TagUseCase.GetByID(ctx, parentID)
		if err != nil {
			return nil, newGraphqlError("AddTag get parent failed", err)
		}
		parent = &model.Tag{
			ID:   model.KeyTag.ToExternalID(parentRet.ID),
			Name: parentRet.Name,
		}
	}
	return &model.AddTagPaylod{
		ClientMutationID: input.ClientMutationID,
		Tag: &model.Tag{
			ID:     model.KeyTag.ToExternalID(useCaseRet.ID),
			Parent: parent,
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

	var parent *model.Tag
	if useCaseRet.ParentTagID > 0 {
		if useCaseRet.ParentTagID > 0 {
			parentRet, err := r.TagUseCase.GetByID(ctx, useCaseRet.ParentTagID)
			if err != nil {
				return nil, newGraphqlError("AddTag get parent failed", err)
			}
			parent = &model.Tag{
				ID:   model.KeyTag.ToExternalID(parentRet.ID),
				Name: parentRet.Name,
			}
		}
	}
	return &model.RemoveTagPayload{
		ClientMutationID: input.ClientMutaionID,
		Tag: &model.Tag{
			ID:     model.KeyTag.ToExternalID(useCaseRet.ID),
			Parent: parent,
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
	useCaseRet, err := r.TagUseCase.GetAll(ctx)
	if err != nil {
		return nil, newGraphqlError("tags operation failed", err)
	}
	tags := map[uint]*model.Tag{}
	for _, v := range useCaseRet {
		tags[v.ID] = &model.Tag{
			ID:   model.KeyTag.ToExternalID(v.ID),
			Name: v.Name,
		}
	}
	// parentをtags内で設定する
	for _, v := range useCaseRet {
		if v.ParentTagID > 0 {
			tags[v.ID].Parent = tags[v.ParentTagID]
		}
	}
	// map to slice
	retTags := make([]*model.Tag, len(tags), len(tags))
	i := 0
	for _, v := range tags {
		retTags[i] = v
		i++
	}
	return retTags, nil
}
