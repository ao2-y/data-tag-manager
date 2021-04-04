package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddItemTemplate(ctx context.Context, input *model.AddItemTemplateInput) (*model.AddItemTemplatePayload, error) {
	itemTemplate, err := r.ItemTemplate.Create(ctx, input.Name, input.MetaKeyIds)
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateItemTemplateMetaKeys(ctx context.Context, input *model.UpdateItemTemplateMetaKeysInput) (*model.UpdateItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveItemTemplate(ctx context.Context, input *model.RemoveItemTemplateInput) (*model.RemoveItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemTemplates(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic(fmt.Errorf("not implemented"))
}
