# Go Backend from scratch
This repository demonstrates how to build a modern backend service in Go by following industry best practices and commonly used architectural patterns.

The project is designed for two main purposes:

- **Interview preparation** – a realistic example of the type of backend service that may be expected during software engineering interviews.
- **Microservice starter** – a clean, production-ready skeleton that can be used as the foundation for new Go microservices.

Rather than focusing on algorithms, this repository emphasizes writing maintainable, scalable, and testable software similar to what is built in production systems.

## What you'll find (something is still in progress)

The project demonstrates many of the practices expected from a modern backend service, including:

- Clean Architecture
- REST and gRPC APIs
- Dependency Injection
- Configuration management
- Structured logging
- Error handling
- Validation
- PostgreSQL integration
- Database migrations
- Unit and integration tests
- Docker & Docker Compose
- Kubernetes deployment
- Graceful shutdown
- Context propagation
- Middleware
- Authentication & authorization
- Health checks
- Observability (metrics, tracing, logging)
- CI/CD-friendly project structure


## Who is this for?

This repository is useful if you want to:

- Prepare for Go backend interviews
- Learn how production Go services are structured
- Explore backend architecture and best practices
- Bootstrap a new microservice instead of starting from scratch
- Use it as a reference implementation for your own projects

The goal is to keep the project practical, easy to understand, and close to real-world backend services used in production.

-----------------------------------
# Get started
## Project structure
- cmd
    - api
        - main.go
- internal
    - app
        - app.go
        - run.go
        - shutdown.go
    - config
        - app.go
    - http
        - router.go
        - handlers/
        - middleware/
    - grpc
        - users
            - handler.go
    - service
        - user_service.go
    - repository
        - user.go
    - domain - core business model and domain-level errors/rules.
        - user.go
        - errors.go
    - db
        - postgres.go
    - observability
        - logger.go
        - metrics.go
        - tracing.go
- migrations
    - sql...
- api
    - proto files
- pkg
- gen

## Engineering Principles
This project follows the same engineering principles used when building production-ready backend services.

- Follow the **[Uber Go Style Guide](https://github.com/uber-go/guide)** and write idiomatic Go
- **Separation of concerns** – keep business logic, transport, persistence, and infrastructure independent.
- **Testability** – write code that is easy to unit and integration test.
- **Graceful shutdown** – properly shut down HTTP/gRPC servers, database connections, background workers, and respect context cancellation.
- **Context propagation** – pass `context.Context` through the entire request lifecycle for cancellation, deadlines, and request-scoped values.
- **Dependency injection** – wire dependencies manually instead of relying on a DI framework to keep dependencies explicit and easy to follow.
- **Composition** over inheritance.
- **Observability**
- **Metrics**
- **Distributed tracing**
- **Clean error handling** – return meaningful domain errors and translate them appropriately at API boundaries.
- **Structured logging with slog**
- **Wrap errors**
- **Health checks**
- **Interfaces where they're consumed** – define interfaces in the package that uses them (typically the service layer). Concrete implementations satisfy them implicitly, reducing coupling and making testing easier.
- **The bigger the interface the weaker abstraction.**
- **Fixed something?** - write a test!
- **SOLID principles** where appropriate
- **Favor simplicity** over clever abstractions

## TODO
- docker-compose
- Makefile
- Dockerfile
- .env.example
- auth middleware
    - store in db

- http to get user

- pipelines + kafka + notification

- makefile/docker
- grpc - generate protobuf. call it as test. README

- graceful shutdown logic
- fail on start

- add metrics / observability
- unit tests
- add integration tests

- validation
- use packages

- tools / flags
- add cache example
- add example of concurrecy! 
