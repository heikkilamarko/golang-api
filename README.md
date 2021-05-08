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

```bash
# Navigate to the kong directory
> cd kong
# Run the configuration script
> ./configure_kong.sh
```

## Run Postman Tests

```bash
# Navigate to the test directory
> cd test
# Run all tests
> ./run_tests.sh
```
