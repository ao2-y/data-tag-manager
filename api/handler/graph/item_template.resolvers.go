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
		id, err := model.KeyMeta.ToInternalID(v)
		if err != nil {
			return nil, fmt.Errorf("invalid MetaKeyID,:%w", err)
		}
		uintMetaKeyIDs[i] = &id
	}
	itemTemplate, err := r.ItemTemplate.Create(ctx, input.Name, uintMetaKeyIDs)
	if err != nil {
		return nil, fmt.Errorf("ItemTemplate Create Failed :%w", err)
	}
	metaKeys := make([]*model.MetaKey, 0, len(itemTemplate.MetaKeys))
	for i, v := range itemTemplate.MetaKeys {
		metaKey := &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(v.ID),
			Name: v.Name,
		}
		metaKeys[i] = metaKey
	}
	retItemTemplate := &model.ItemTemplate{
		ID:       model.KeyItemTemplate.ToExternalID(itemTemplate.ID),
		Name:     itemTemplate.Name,
		MetaKeys: metaKeys,
	}

	return &model.AddItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     retItemTemplate,
	}, nil
}

func (r *mutationResolver) UpdateItemTemplateName(ctx context.Context, input *model.UpdateItemTemplateNameInput) (*model.UpdateItemTemplatePayload, error) {
	uintId, err := model.KeyItemTemplate.ToInternalID(input.ItemTemplateID)
	if err != nil {
		return nil, newGraphqlError("ID format error", err)
	}
	ret, err := r.ItemTemplate.UpdateName(ctx, uintId, input.Name)
	if err != nil {
		return nil, newGraphqlError("UpdateName failed", err)
	}
	metaKeys := make([]*model.MetaKey, len(ret.MetaKeys), len(ret.MetaKeys))
	for i, v := range ret.MetaKeys {
		metaKey := &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(v.ID),
			Name: v.Name,
		}
		metaKeys[i] = metaKey
	}
	retItemTemplate := &model.ItemTemplate{
		ID:       model.KeyItemTemplate.ToExternalID(ret.ID),
		Name:     ret.Name,
		MetaKeys: metaKeys,
	}
	return &model.UpdateItemTemplatePayload{
		ClientMutationID: input.ClientMutationID,
		ItemTemplate:     retItemTemplate,
	}, nil
}

func (r *mutationResolver) UpdateItemTemplateMetaKeys(ctx context.Context, input *model.UpdateItemTemplateMetaKeysInput) (*model.UpdateItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveItemTemplate(ctx context.Context, input *model.RemoveItemTemplateInput) (*model.RemoveItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemTemplates(ctx context.Context) ([]*model.ItemTemplate, error) {
	its, err := r.ItemTemplate.FetchAll(ctx)
	if err != nil {
		return nil, newGraphqlError("ItemTemplates failed", err)
	}
	ret := make([]*model.ItemTemplate, len(its), len(its))
	for i, v := range its {
		metaKeys := make([]*model.MetaKey, len(v.MetaKeys), len(v.MetaKeys))
		for i2, v2 := range v.MetaKeys {
			metaKeys[i2] = &model.MetaKey{
				ID:   model.KeyMeta.ToExternalID(v2.ID),
				Name: v2.Name,
			}
		}
		ret[i] = &model.ItemTemplate{
			ID:       model.KeyItemTemplate.ToExternalID(v.ID),
			Name:     v.Name,
			MetaKeys: metaKeys,
		}
	}
	return ret, nil
}
