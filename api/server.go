package main

import (
	"ao2-y/data-tag-manager/handler/graph"
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/infra/persistent/mysql"
	"ao2-y/data-tag-manager/usecase"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ItemUseCase: usecase.NewItemUseCase(),
		ItemTemplate: usecase.NewItemTemplateUseCase(
			mysql.NewItemTemplateRepository(mysql.NewDBConnection(
				"localhost",
				"3306",
				"admin",
				"password",
				"data_tag_manager",
			))),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
