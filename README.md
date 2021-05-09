# A Simple RESTful API Written in Go

## Tech Stack

- [Go](https://go.dev/)
- [OpenAPI 3](https://www.openapis.org/)
- [Kong Gateway](https://konghq.com/kong/)
- [PostgreSQL](https://www.postgresql.org/)
- [Migrate](https://github.com/golang-migrate/migrate)
- [Postman](https://www.postman.com/) & [Newman](https://www.npmjs.com/package/newman)
- [Docker](https://www.docker.com/)

## Build and Run

```bash
# Build and run
> docker compose up --build
```

## Configure Kong Gateway

### _Option 1:_ curl

```bash
# Navigate to the kong/curl directory
> cd kong/curl
# Run the configuration script
> ./configure_kong.sh
```

### _Option 2:_ postman/newman

```bash
# Navigate to the kong/postman directory
> cd kong/postman
# Run the configuration script
> ./configure_kong.sh
```

### _Option 3:_ Kong Manager

Kong Manager is the visual, browser-based tool for monitoring and managing Kong Gateway.

Kong Manager URL: http://localhost:8002

### _Option 4:_ DB-less and Declarative Configuration

[Kong Docs](https://docs.konghq.com/gateway-oss/2.4.x/db-less-and-declarative-config/)

See `kong/declarative/kong.yml`

## Run Postman Tests

```bash
# Navigate to the test directory
> cd test
# Run all tests
> ./run_tests.sh
```
