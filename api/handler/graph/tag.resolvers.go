package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddTag(ctx context.Context, input *model.AddTagInput) (*model.AddTagPaylod, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveTag(ctx context.Context, input *model.RemoveTagInput) (*model.RemoveTagPayload, error) {
	panic(fmt.Errorf("not implemented"))
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
