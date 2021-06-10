package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddItemTemplate(ctx context.Context, input *model.AddItemTemplateInput) (*model.AddItemTemplatePayload, error) {
	uintMetaKeyIDs := make([]*uint, len(input.MetaKeyIds))
	for i, v := range input.MetaKeyIds {
		id, err := model.IDTypeMeta.ToInternalID(v)
		if err != nil {
			return nil, fmt.Errorf("invalid MetaKeyID,:%w", err)
		}
		uintMetaKeyIDs[i] = &id
	}
	itemTemplate, err := r.ItemTemplate.Create(ctx, input.Name, uintMetaKeyIDs)
	if err != nil {
		return nil, fmt.Errorf("ItemTemplate Create Failed :%w", err)
	}

	return &model.AddItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     model.NewItemTemplate(itemTemplate),
	}, nil
}

func (r *mutationResolver) UpdateItemTemplateName(ctx context.Context, input *model.UpdateItemTemplateNameInput) (*model.UpdateItemTemplatePayload, error) {
	uintId, err := model.IDTypeItemTemplate.ToInternalID(input.ItemTemplateID)
	if err != nil {
		return nil, newGraphqlError("ID format error", err)
	}
	ret, err := r.ItemTemplate.UpdateName(ctx, uintId, input.Name)
	if err != nil {
		return nil, newGraphqlError("UpdateName failed", err)
	}

	return &model.UpdateItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     model.NewItemTemplate(ret),
	}, nil
}

func (r *mutationResolver) UpdateItemTemplateMetaKeys(ctx context.Context, input *model.UpdateItemTemplateMetaKeysInput) (*model.UpdateItemTemplatePayload, error) {
	uintId, err := model.IDTypeItemTemplate.ToInternalID(input.ItemTemplateID)
	if err != nil {
		return nil, newGraphqlError("ID format error", err)
	}
	uintMetaKeyIDs := make([]*uint, len(input.MetaKeyIds), len(input.MetaKeyIds))
	for i, v := range input.MetaKeyIds {
		uid, err := model.IDTypeMeta.ToInternalID(v)
		if err != nil {
			return nil, newGraphqlError("ID format error", err)
		}
		uintMetaKeyIDs[i] = &uid
	}
	ret, err := r.ItemTemplate.UpdateMetaKeys(ctx, uintId, uintMetaKeyIDs)
	return &model.UpdateItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     model.NewItemTemplate(ret),
	}, nil
}

func (r *mutationResolver) RemoveItemTemplate(ctx context.Context, input *model.RemoveItemTemplateInput) (*model.RemoveItemTemplatePayload, error) {
	uintID, err := model.IDTypeItemTemplate.ToInternalID(input.ItemTemplateID)
	if err != nil {
		return nil, newGraphqlError("ID format error", err)
	}
	it, err := r.ItemTemplate.Remove(ctx, uintID)
	if err != nil {
		return nil, newGraphqlError("RemoveItemTemplate failed.", err)
	}

	return &model.RemoveItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     model.NewItemTemplate(it),
	}, nil
}

func (r *queryResolver) ItemTemplates(ctx context.Context) ([]*model.ItemTemplate, error) {
	its, err := r.ItemTemplate.FetchAll(ctx)
	if err != nil {
		return nil, newGraphqlError("ItemTemplates failed", err)
	}

	return model.NewItemTemplates(its), nil
}
