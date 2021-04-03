package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddItemTemplate(ctx context.Context, input *model.AddItemTemplateInput) (*model.AddItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateItemTemplateName(ctx context.Context, input *model.UpdateItemTemplateNameInput) (*model.UpdateItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateItemTemplateMetaKeys(ctx context.Context, input *model.UpdateItemTemplateMetaKeysInput) (*model.UpdateItemTemplatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemTemplates(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic(fmt.Errorf("not implemented"))
}
