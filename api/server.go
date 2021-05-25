package main

import (
	"ao2-y/data-tag-manager/handler/graph/generated"
	"ao2-y/data-tag-manager/injector"
	"ao2-y/data-tag-manager/logger"
	"ao2-y/data-tag-manager/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	applogger := logger.InitApplicationLogger()
	router.Use(middleware.ContextLogger(applogger))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(injector.NewGraphqlConfig(applogger)))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
