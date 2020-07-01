# Products API

## API

- Golang
- OpenAPI 3.0 (design-first with Swagger Editor)
- Swagger UI (static website hosting in Azure Storage)

## Database

- Azure PostgreSQL Database
- Flyway (database migrations)

## Local Development Setup

### Create Directory for Secrets

```bash
> mkdir ../golang-api-secrets
```

Creating the folder outside of the repository prevents config being pushed into the repo accidentally.

### Create Files

Create the following files into the folder created in the previous step.

```bash
# api.env

APP_API_KEY=
APP_PORT=
APP_DB_HOST=
APP_DB_PORT=
APP_DB_NAME=
APP_DB_USERNAME=
APP_DB_PASSWORD=
APP_CORS_ENABLED=
```

```bash
# acr.env

ACR_SERVER=
ACR_USER=
ACR_PASSWORD=
```

```bash
# flyway.env

FLYWAY_URL=
FLYWAY_USER=
FLYWAY_PASSWORD=
FLYWAY_SCHEMAS=
```

```bash
# pgadmin.env

PGADMIN_DEFAULT_EMAIL=
PGADMIN_DEFAULT_PASSWORD=
```

```bash
# postman_environment.json
```
