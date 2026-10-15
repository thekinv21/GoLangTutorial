**English** | [Turkish](./ROADMAP_TR.md) | [Russian](./ROADMAP_RU.md)

# Go (Golang) Learning Roadmap

## Stage 0 — Installation and Introduction

- [x] What is Go? ([README_EN.md](../01_installation/docs/README_EN.md))
- [x] Installing Go
- [x] VS Code + Go extension
- [x] First program: `Hello, World!`
- [x] `go run`, `go build`, `go fmt`, `go vet`
- [x] `go mod init` and `go.mod`
- [x] `package main` and `func main()`

**Exercises:**

- You can write a program that prints your name and age.
- You can compile the same program with `go build` and run the resulting file.

## Stage 1 — Basics

- [x] Variables ([README_EN.md](../02_basics/docs/README_EN.md))
- [x] Constants
- [x] Basic types
- [x] Type conversion
- [ ] Operators
- [ ] The `fmt` package
- [ ] Reading user input
- [ ] Conditions: `if / else`, `switch`
- [ ] Loops: `for`

**Exercises:**

- You can build a command-line calculator.
- You can build a number guessing game.
- You can print FizzBuzz from 1 to 100.

## Stage 2 — Functions and Data Structures

- [ ] Functions
- [ ] Multiple return values
- [ ] Variadic functions
- [ ] Anonymous functions and closures
- [ ] `defer`
- [ ] Arrays
- [ ] Slices
- [ ] Maps
- [ ] String operations
- [ ] The `strconv` package

**Exercises:**

- You can write a program that counts how many times each word appears in a text.
- You can build a grade average calculator.

## Stage 3 — Pointers, Structs and Methods

- [ ] Pointers
- [ ] Pass by value vs pass by pointer
- [ ] Structs
- [ ] Struct embedding
- [ ] Methods
- [ ] Constructor pattern
- [ ] Struct tags

**Exercises:**

- You can build an in-memory library or student management system.

## Stage 4 — Interfaces and Error Handling

- [ ] Interfaces
- [ ] `any` / `interface{}`
- [ ] Type assertion and type switch
- [ ] Common interfaces
- [ ] Error handling
- [ ] Custom error types
- [ ] `errors.Is`, `errors.As`
- [ ] `panic` and `recover`

**Exercises:**

- You can write a program that calculates shape areas using a `Shape` interface.
- You can build a bank account that returns an error on invalid operations.

## Stage 5 — Packages, Modules and Project Structure

- [ ] Writing your own package
- [ ] Adding external packages
- [ ] Project folder structure
- [ ] The `init()` function

**Exercises:**

- You can split one of your previous projects into packages.

## Stage 6 — Generics

- [ ] Generic functions
- [ ] Type constraints
- [ ] Generic structs
- [ ] The `slices` and `maps` packages

**Exercises:**

- You can write a generic Stack or Queue.

## Stage 7 — Concurrency

- [ ] Goroutines
- [ ] `sync.WaitGroup`
- [ ] Channels
- [ ] `select`
- [ ] Mutex
- [ ] Race conditions
- [ ] The `context` package
- [ ] Concurrency patterns

**Exercises:**

- You can build a tool that checks a list of URLs concurrently and prints their status codes.
- You can write a program that processes files with a worker pool.

## Stage 8 — Testing and Code Quality

- [ ] The `testing` package
- [ ] Table-driven tests
- [ ] Benchmarks
- [ ] Test coverage
- [ ] Mocking
- [ ] `golangci-lint`

**Exercises:**

- You can write tests for your previous projects.

## Stage 9 — Standard Library

- [ ] Reading / writing files
- [ ] JSON
- [ ] The `time` package
- [ ] Command-line arguments
- [ ] Logging
- [ ] Environment variables

**Exercises:**

- You can build a CLI Todo app that stores tasks in a JSON file.

## Stage 10 — Web and Backend

- [ ] HTTP server
- [ ] Routing
- [ ] JSON API
- [ ] Middleware
- [ ] HTTP client
- [ ] Frameworks
- [ ] Database
- [ ] ORM / query builder
- [ ] Migrations
- [ ] Config management
- [ ] Authentication
- [ ] Validation

**Exercises:**

- You can build a Blog or Todo REST API with PostgreSQL and JWT.

## Stage 11 — Advanced Topics and Deployment

- [ ] Clean / Hexagonal Architecture
- [ ] Graceful shutdown
- [ ] Docker
- [ ] Docker Compose
- [ ] gRPC
- [ ] WebSocket
- [ ] Redis
- [ ] Message queues
- [ ] Profiling
- [ ] CI/CD

**Exercises:**

- You can build a microservice-based URL shortener, chat app or e-commerce backend.
