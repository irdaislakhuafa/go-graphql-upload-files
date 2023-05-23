package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/domain"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/bussines/usecase"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/config"
	"github.com/irdaislakhuafa/go-graphql-upload-files/src/schema/graph"
)

const (
	configFilePath = "etc/config/config.jsonc"
)

func main() {
	ctx := context.Background()

	// read config
	config, err := config.Init(ctx, configFilePath)
	if err != nil {
		panic(err)
	}

	// init domain
	domain := domain.Init(ctx, config)

	// init usecase
	usecase := usecase.Init(ctx, config, *domain)

	// gql config
	gqlConfig := graph.Config{
		Resolvers: &graph.Resolver{
			Todo: usecase.Todo,
		},
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(gqlConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Application.Port)
	log.Fatal(http.ListenAndServe(":"+config.Application.Port, nil))
}
