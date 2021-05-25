package injector

import (
	"ao2-y/data-tag-manager/handler/graph"
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/infra/persistent/mysql"
	"ao2-y/data-tag-manager/usecase"
	"go.uber.org/zap"
)

func NewGraphqlConfig(logger *zap.Logger) generated.Config {
	config := newConfig()
	dbCon := mysql.NewDBConnection(
		logger,
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DatabaseName,
	)
	metaRepository := mysql.NewMetaRepository(dbCon, true)
	itemTemplateRepository := mysql.NewItemTemplateRepository(dbCon)
	itemUseCase := usecase.NewItemUseCase()
	itemTemplateUseCase := usecase.NewItemTemplateUseCase(itemTemplateRepository, metaRepository)
	metaUseCase := usecase.NewMetaUseCase(metaRepository)
	tagRepository := mysql.NewTagRepository(dbCon, true)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	return generated.Config{
		Resolvers: &graph.Resolver{
			ItemUseCase:  itemUseCase,
			ItemTemplate: itemTemplateUseCase,
			MetaUseCase:  metaUseCase,
			TagUseCase:   tagUseCase,
		},
	}
}
