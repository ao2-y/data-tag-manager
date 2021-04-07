package graph

import "ao2-y/data-tag-manager/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ItemUseCase  usecase.Item
	ItemTemplate usecase.ItemTemplate
	MetaUseCase  usecase.Meta
}
