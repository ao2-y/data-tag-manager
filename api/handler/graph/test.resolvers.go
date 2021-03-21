package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddItem(ctx context.Context, item model.NewItem) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Items(ctx context.Context, ids []*string) ([]*model.Item, error) {
	tmp := uint(1)
	dummy := []*uint{&tmp}
	usecaseRet, err := r.ItemUseCase.GetItems(ctx, dummy)
	if err != nil {
		return nil, err
	}

	var ret []*model.Item
	for _, v := range usecaseRet {
		item := &model.Item{
			ID:          fmt.Sprintf("%v", v.ID),
			Name:        v.Name,
			Description: v.Description,
			Metas:       nil,
			Tags:        nil,
		}
		ret = append(ret, item)
	}
	return ret, nil
}

func (r *queryResolver) ItemTemplates(ctx context.Context) ([]*model.ItemTemplate, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
