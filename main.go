package main

import (
	"fmt"
	"net/http"

	"github.com/graph-gophers/graphql-go"
)

// Definir el esquema de GraphQL
const schemaString = `
    type Query {
        hello: String!
    }
`

// Resolver que manejará la consulta 'hello'
type Resolver struct{}

func (r *Resolver) Hello() string {
	return "Hello, World!"
}

func main() {
	// Crear el esquema de GraphQL
	schema := graphql.MustParseSchema(schemaString, &Resolver{})

	// Crear un handler para GraphQL
	http.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ejecutar la consulta GraphQL
		schema.Exec(r.Context(), r.URL.Query().Get("query"))
	}))

	// Configurar el servidor HTTP
	fmt.Println("Servidor corriendo en http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
