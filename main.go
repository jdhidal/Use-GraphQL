package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

const schemaString = `
    type Query {
        hello: String!
        currentTime: String!
    }
`

// Resolver
type Resolver struct{}

// Method handling the "hello" query
func (r *Resolver) Hello() string {
	// Print in the terminal when the "hello" query is made
	fmt.Println("Request made: 'hello' query received")
	return "Hello, World!"
}

// Method handling the "currentTime" query
func (r *Resolver) CurrentTime() string {
	// Get the current time
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// Print in the terminal when the "currentTime" query is made
	fmt.Println("Request made: 'currentTime' query received")
	return currentTime
}

func main() {
	// Create the GraphQL schema
	schema := graphql.MustParseSchema(schemaString, &Resolver{})

	// Configure the HTTP handler
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	// Start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
