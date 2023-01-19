package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/C-Benzz/assignment1.git/graph"
	"github.com/C-Benzz/assignment1.git/graph/model"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: graph.Database{
			TodoTable: make(map[string]model.Todo),
		},
	}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
