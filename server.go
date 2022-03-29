package main

import (
	"log"
	"net/http"
	"os"

	"github.com/harryduong99/songchord-api-v2/driver"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/harryduong99/songchord-api-v2/graph"
	"github.com/harryduong99/songchord-api-v2/graph/generated"
)

const defaultPort = "8080"

func main() {
	loadEnv()
	connectDb()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func connectDb() {
	driver.ConnectDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
