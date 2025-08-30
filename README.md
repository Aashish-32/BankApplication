# Bank API

This project implements a simple bank API using Go, PostgreSQL, and Docker. It provides functionalities for managing user accounts, transfers, and authentication.

## Features

*   User authentication (signup, login)
*   Account management (create, get, list, update, delete)
*   Money transfers between accounts
*   Secure password hashing
*   JWT and PASETO token-based authentication
*   Database migrations

## Technologies Used

*   **Go (Golang):** Backend API development.
*   **PostgreSQL:** Relational database for storing bank data.
*   **SQLC:** Generates Go code from SQL queries, ensuring type safety and reducing boilerplate.
*   **Docker & Docker Compose:** Containerization and orchestration for easy setup and deployment.
*   **Gin Web Framework:** For building the RESTful API.
*   **`github.com/golang-migrate/migrate`:** For database schema migrations.
*   **`github.com/lib/pq`:** PostgreSQL driver for Go.
*   **`github.com/joho/godotenv`:** For loading environment variables from `.env` file.
*   **PASETO & JWT:** For secure token-based authentication.
*   **GitHub Actions:** For CI/CD (as indicated by `.github/workflows/ci.yml`).

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:

*   [Go](https://golang.org/doc/install) (version 1.20 or higher recommended)
*   [Docker](https://docs.docker.com/get-docker/) (Docker Desktop for Windows/macOS, or Docker Engine for Linux)
*   [Docker Compose](https://docs.docker.com/compose/install/)
*   [`migrate` tool](https://github.com/golang-migrate/migrate): Used for database migrations. You can install it via:
    ```bash
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Aashish-32/bank.git
    cd bank
    ```

2.  **Create a `.env` file:**
    Copy the `.env.example` (if it exists, otherwise create it) and populate it with your environment variables. A minimal `.env` file should look like this:
    ```
    DB_DRIVER=postgres
    DB_SOURCE="postgresql://root:password@localhost:5432/simplebank?sslmode=disable"
    SERVER_ADDRESS=0.0.0.0:8080
    TOKEN_SYMMETRIC_KEY=your_secret_key_at_least_32_chars_long
    ```
    *Note: The `SERVER_ADDRESS` in `main.go` and `Dockerfile` uses `8080`, while `docker-compose.yaml` uses `8000`. For consistency, we'll use `8080` in the `.env` and `Makefile` examples, but be aware of this discrepancy if running with Docker Compose.* For Docker Compose, the `api` service exposes port `8000` internally and maps it to `8000` externally.

### Running the Application

There are two primary ways to run the application: using Docker Compose (recommended for development) or running it locally.

#### Using Docker Compose (Recommended)

This method sets up both the PostgreSQL database and the Go API server in isolated Docker containers.

1.  **Build and run the services:**
    ```bash
    docker-compose up --build
    ```
    This command will:
    *   Build the `api` service Docker image.
    *   Start the `postgres` service.
    *   Start the `api` service, waiting for the `postgres` service to be ready.
    *   Run database migrations automatically via `start.sh`.

2.  **Access the API:**
    The API server will be accessible at `http://localhost:8000` (as configured in `docker-compose.yaml`).

3.  **Stop the services:**
    To stop and remove the containers, networks, and volumes created by `docker-compose up`:
    ```bash
    docker-compose down
    ```

#### Running Locally

To run the application directly on your machine, you'll need a running PostgreSQL instance.

1.  **Start PostgreSQL (if not already running):**
    You can use Docker to run a PostgreSQL container:
    ```bash
    docker run --name mypostgres --network bank-network -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -p 5432:5432 -d postgres:16.0-alpine3.18
    ```
    *Note: If you don't have `bank-network`, you might need to create it or remove `--network bank-network`.*

2.  **Create the database:**
    ```bash
    docker exec -it mypostgres createdb --username=root --owner=root simplebank
    ```

3.  **Run database migrations:**
    Ensure the `migrate` tool is installed (see Prerequisites).
    ```bash
    migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up
    ```

4.  **Run the Go application:**
    ```bash
    go run main.go
    ```
    The API server will be accessible at `http://localhost:8080` (as configured in `main.go` and `.env`).

## Database Setup and Migrations

The project uses `golang-migrate` for database migrations. Migration files are located in the `db/migration` directory.

*   **Apply all pending migrations:**
    ```bash
    make migrateup
    ```
    or directly:
    ```bash
    migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up
    ```

*   **Rollback the last migration:**
    ```bash
    make migratedown1
    ```
    or directly:
    ```bash
    migrate -path db/migration -database "postgresql://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down 1
    ```

*   **Create a new migration:**
    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```
    Replace `<migration_name>` with a descriptive name for your migration.

## API Endpoints

The API endpoints are defined in the `api` directory. You can explore `api/server.go`, `api/users.go`, `api/accounts.go`, and `api/transfer.go` to understand the available routes and their functionalities.

## Testing

To run the unit and integration tests for the project:

```bash
make test
```

or directly:

```bash
go test -v -cover ./...
```

## Generating SQLC Code and Mocks

*   **Generate SQLC code:**
    If you modify SQL queries in `db/query`, you'll need to regenerate the Go code:
    ```bash
    make sqlc
    ```

*   **Generate Mocks:**
    To generate mock interfaces for testing:
    ```bash
    make mock
    ```

