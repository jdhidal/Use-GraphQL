# GraphQL with Go

This project demonstrates how to build a simple GraphQL server using Go and the `graphql-go` library. The server has two queries: `hello` and `currentTime`. The goal of this project is to showcase how to integrate GraphQL in a Go-based backend service.

## Why Go and GraphQL?

Go is a statically typed, compiled programming language that is known for its simplicity and performance. It is ideal for building scalable and high-performance services. GraphQL, on the other hand, is a query language for APIs that allows clients to request only the data they need. This makes it a powerful choice for building flexible APIs with Go.

In this project, we use Go to create a simple GraphQL server, and GraphQL to handle two queries (`hello` and `currentTime`). This demonstrates how Go can be used to implement efficient and flexible GraphQL APIs.

## What the Code Does

The code implements a basic GraphQL server with the following features:
- **Schema**: Defines two queries: `hello` and `currentTime`.
- **Resolver**: Implements the logic behind each query.
  - `hello`: Returns the string "Hello, World!".
  - `currentTime`: Returns the current date and time in the format `YYYY-MM-DD HH:MM:SS`.
- **Server**: A simple HTTP server that listens on port 8080 and exposes a `/graphql` endpoint.

### Queries:
1. **hello**: Returns "Hello, World!".
2. **currentTime**: Returns the current date and time.

## How to Run the Code

### Prerequisites

- Install Go (version 1.16 or higher). If you don't have Go installed, you can download it from [https://golang.org/dl/](https://golang.org/dl/).
- Install the `graphql-go` package by running:
  ```bash
  go get github.com/graph-gophers/graphql-go

### Steps to Run
1. **Clone the repository to your local machine:**
    ```bash
    git clone https://github.com/jdhidal/Use-GraphQL.git
    cd Use-Webhook
    ```

2. **Run the Go application:**
    ```bash
    go get github.com/graph-gophers/graphql-go
    go run main.go

3. **The server will start and listen on port 8080. You can test it by sending a POST request to http://localhost:8080/graphql with the following queries:**

    - Query hello:
        ```bash
        curl -X POST -H "Content-Type: application/json" -d "{\"query\": \"{ hello }\"}" http://localhost:8080/graphql

    - Query currentTime:
        ```bash
        curl -X POST -H "Content-Type: application/json" -d "{\"query\": \"{ currentTime }\"}" http://localhost:8080/graphql

### Example Response
- **hello**: {"data":{"hello":"Hello, World!"}}
- **currentTime**: {"data":{"currentTime":"2023-03-01 14:30: