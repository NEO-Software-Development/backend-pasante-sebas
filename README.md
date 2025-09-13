# Go Task Manager API

This is a simple RESTful API for a task manager application, written in Go. It's designed as a microservice and can be easily deployed as a proof of concept.

## Features

*   **RESTful API:** Provides endpoints for CRUD (Create, Read, Update, Delete) operations on tasks.
*   **Flexible Database:** Uses GORM for object-relational mapping, with support for SQLite (for development) and PostgreSQL (for production).
*   **Auto-generated Documentation:** API documentation is automatically generated using Swagger/OpenAPI.

## API Documentation

The API documentation is available at the `/swagger/index.html` endpoint. For example, if you are running the application locally, you can access the documentation at `http://localhost:8080/swagger/index.html`.

## API Endpoints

All endpoints are prefixed with `/api/v1`.

| Method | Endpoint      | Description          |
|--------|---------------|----------------------|
| POST   | `/tasks`      | Create a new task    |
| GET    | `/tasks`      | Get all tasks        |
| GET    | `/tasks/{id}` | Get a task by ID     |
| PUT    | `/tasks/{id}` | Update a task        |
| DELETE | `/tasks/{id}` | Delete a task        |

## Setup

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/backend-pasante-sebas.git
    cd backend-pasante-sebas
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Configure the environment:**

    The application can be configured using environment variables. The following variables are available:

    *   `DB_DRIVER`: The database driver to use. Can be `sqlite` or `postgres`. Defaults to `sqlite`.
    *   `DSN`: The data source name for the database connection. Defaults to `tasks.db` for SQLite. For PostgreSQL, it should be a connection string like `host=localhost user=user password=password dbname=tasks port=5432 sslmode=disable`.
    *   `PORT`: The port on which to run the server. Defaults to `8080`.

## Usage

1.  **Run the application:**

    ```bash
    go run cmd/api/main.go
    ```

2.  **Access the API:**

    You can use a tool like `curl` or Postman to interact with the API.

    **Example: Create a new task**

    ```bash
    curl -X POST http://localhost:8080/api/v1/tasks \
    -H "Content-Type: application/json" \
    -d '{"name": "My new task", "description": "This is a new task"}'
    ```
