package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/handler/graph/model"
	"ao2-y/data-tag-manager/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
)

func (r *mutationResolver) Noop(ctx context.Context, input *model.NoopInput) (*model.NoopPayload, error) {
	return &model.NoopPayload{ClientMutationID: input.ClientMutationID}, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	logger.Ctx(ctx).Debug("Test!!")
	uintId, keyType, err := model.IDtoKeyNameAndInternalID(id)
	ctx = logger.With(ctx, zap.String("ID", id))
	if err != nil {
		return nil, fmt.Errorf("ID type error:%w", err)
	}
	switch keyType {
	case model.KeyUnknown:
		return nil, fmt.Errorf("ID type unkown")
	case model.KeyItemTemplate:
		it, err := r.ItemTemplate.FetchByID(ctx, uintId)
		if err != nil {
			return nil, fmt.Errorf("ItemTemplate FetchByID Failed.%w", err)
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
		ret, err := r.MetaUseCase.FetchKeyByID(ctx, uintId)
		if err != nil {
			return nil, newGraphqlError("", err)
		}
		return &model.MetaKey{
			ID:   model.KeyMeta.ToExternalID(ret.ID),
			Name: ret.Name,
		}, nil
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
