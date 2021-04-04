package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/handler/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) Noop(ctx context.Context, input *model.NoopInput) (*model.NoopPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	uintId, keyType, err := model.IDtoKeyNameAndInternalID(id)
	if err != nil {
		return nil, fmt.Errorf("ID type error:%w", err)
	}
	switch keyType {
	case model.KeyUnknown:
		return nil, fmt.Errorf("ID type unkown")
	case model.KeyItemTemplate:
		it, err := r.ItemTemplate.Fetch(ctx, uintId)
		if err != nil {
			return nil, fmt.Errorf("ItemTemplate Fetch Failed.%w", err)
		}
		return &model.ItemTemplate{
			ID:       model.KeyItemTemplate.ToExternalID(it.ID),
			Name:     it.Name,
			MetaKeys: nil, // TODO
		}, nil
	case model.KeyItem:
		return nil, fmt.Errorf("item not implement")
	case model.KeyTag:
		return nil, fmt.Errorf("tag not implement")
	case model.KeyMeta:
		return nil, fmt.Errorf("meta not implement")
	default:
		panic("ここにくることがない")
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
