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
> docker compose up --build
```

Kong Manager: http://localhost:8002

## Run Postman Tests

```bash
# Navigate to the test directory
> cd test
# Run all tests
> ./run_tests.sh
```
