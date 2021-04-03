package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddItem(ctx context.Context, input *model.AddItemInput) (*model.AddItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveItem(ctx context.Context, input *model.RemoveItemInput) (*model.RemoveItemPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context, first *int, after *string, last *int, before *string) (*model.ItemConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
