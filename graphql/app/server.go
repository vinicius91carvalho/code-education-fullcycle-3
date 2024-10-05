package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vinicius91carvalho/code-education-graphql/graph"
	"github.com/vinicius91carvalho/code-education-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	log.Print("Starting server...")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v",err)
	}
	defer db.Close()
	log.Print("Database connection established")

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)
	log.Print("Database handler initialized")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Using port: %s", port)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB: courseDb,
	}}))
	log.Print("GraphQL server created")

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
