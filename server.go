package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lakshyab1995/tiger-kittens/auth"
	"github.com/lakshyab1995/tiger-kittens/db"
	"github.com/lakshyab1995/tiger-kittens/graph"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env") // Load environment variables from .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	log.Println(port)
	if port == "" {
		port = defaultPort
	}

	gDB, err := db.Connect() // Connect to the database
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	resolver := graph.NewResolver(gDB) // Create a new resolver

	router := chi.NewRouter()

	router.Use(auth.Middleware(resolver))

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
