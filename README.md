# TODO
1. Preparing
## Infra
- docker-compose
- Makefile
- Dockerfile
- .env.example

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

## Approach
- separate concerns
- test everything
- graceful shutdown - http, db, cancellation
- metrics / observability
- context
- wire dependencies manually
- Observability
  - slog
  - tracing
+ interfaces - define the interface where it's consumed (service package) and SomeRepository satisfies it implicitly.

2. To work on
- auth middleware
    - store in db

- http to get user

- pipelines + kafka + notification

- makefile/docker
- grpc - generate protobuf. call it as test. README

- graceful shutdown logic
- fail on start

- metrics / observability
- tests
- integration tests

- validation
- packages

- tools / flags