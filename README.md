# A Simple RESTful API Written in Go

## Tech Stack

- [Go](https://go.dev/)
- [OpenAPI 3](https://www.openapis.org/)
- [Kong Gateway](https://konghq.com/kong/)
- [PostgreSQL](https://www.postgresql.org/)
- [Migrate](https://github.com/golang-migrate/migrate)
- [Postman](https://www.postman.com/) & [Newman](https://www.npmjs.com/package/newman)
- [Docker](https://www.docker.com/)

## Secrets Management

This repository uses [SOPS](https://github.com/mozilla/sops) with [age](https://github.com/mozilla/sops#encrypting-using-age) for managing secrets.

See [secrets](secrets/) for details.

## Build and Run

Before running the below command, make sure you have the unencrypted secrets in the `env` directory. See [secrets](secrets/) for details.

```bash
> docker compose up --build -d
```

Kong Manager: http://localhost:8002

## Run Postman Tests

```bash
# Navigate to the test directory
> cd test
# Run tests
> ./run_tests.sh
```
