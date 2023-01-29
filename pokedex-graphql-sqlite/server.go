package main

import (
	"log"
	"net/http"
	"pokedex-graphql/database"
	"pokedex-graphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	// establish connection
	database.ConnectDB()
	// create db
	database.CreateDB()
	// migrate the db with Pokemon model
	database.MigrateDB()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: database.DBInstance}}))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
