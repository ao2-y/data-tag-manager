package injector

import (
	"ao2-y/data-tag-manager/handler/graph"
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/infra/persistent/mysql"
	"ao2-y/data-tag-manager/usecase"
)

func NewGraphqlConfig() generated.Config {
	config := newConfig()
	dbCon := mysql.NewDBConnection(
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DatabaseName,
	)
	metaRepository := mysql.NewMetaRepository(dbCon)
	itemTemplateRepository := mysql.NewItemTemplateRepository(dbCon)
	itemUseCase := usecase.NewItemUseCase()
	itemTemplateUseCase := usecase.NewItemTemplateUseCase(itemTemplateRepository, metaRepository)
	metaUseCase := usecase.NewMetaUseCase(metaRepository)
	return generated.Config{
		Resolvers: &graph.Resolver{
			ItemUseCase:  itemUseCase,
			ItemTemplate: itemTemplateUseCase,
			MetaUseCase:  metaUseCase,
		},
	}
}
