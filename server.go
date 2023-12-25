package main

import (
	"log"
	"net/http"
	"os"
	"queue/config"
	"queue/database"
	"queue/graph"
	"queue/usecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conf := config.New("./config")

	db := database.New(&conf.DB)
	usc := usecase.New(db, &conf.Usecase)
	go usc.Batch()

	resolver := graph.NewResolver(usc)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
