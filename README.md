# CityWalker API

## Getting Started

## MakeFile
run all tests
```bash
make test
```
lint the code
```bash
make lint
```
run the application in production mode
```bash
make run-prod
```
run the application in development mode
```bash
make run-dev
```

## ENVIRONMENT
```
# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS=

ENV_FILE=

# Server settings:
SERVER_HOST=
SERVER_PORT=
SERVER_READ_TIMEOUT=
PREFORK=

# Database settings:
DATABASE=
DB_HOST=
DB_PORT=

# MongoDB settings:
MDB_DATABASE=
MDB_USERNAME=
MDB_PASSWORD=
MDB_COLLECTION_USERS=
MDB_COLLECTION_CITIES=
MDB_COLLECTION_TRAVELS=


# JWT
JWT_SECRET_KEY=
JWT_EXPIRE_HOUR_COUNT=

# SMTP
#EMAIL
SMTP_SERVER=
SMTP_PORT=
SMTP_USERNAME=
SMTP_PASSWORD=
```