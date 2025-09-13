# Go Task Manager

This is a simple command-line task manager written in Go. It allows you to manage a list of tasks stored in a MySQL database.

## Features

*   View all tasks
*   Add a new task
*   Mark a task as completed
*   Delete a task

## Prerequisites

*   Go (version 1.15 or later)
*   MySQL

## Setup

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/backend-pasante-sebas.git
    cd backend-pasante-sebas
    ```

2.  **Set up the database:**

    *   Make sure you have MySQL installed and running.
    *   Create a database named `tasks`.
    *   Create a table named `task_list` with the following schema:

    ```sql
    CREATE TABLE task_list (
        id INT AUTO_INCREMENT PRIMARY KEY,
        tasks VARCHAR(255) NOT NULL,
        description TEXT,
        completed BOOLEAN NOT NULL DEFAULT FALSE
    );
    ```

3.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

## Usage

1.  **Run the application:**

    ```bash
    go run main.go
    ```

2.  **Follow the on-screen menu:**

    The application will present you with a menu of options to manage your tasks. Simply enter the number corresponding to the action you want to perform.

    ```
    Selecciona la opción deseada:
    1. Ver tareas
    2. Agregar tarea
    3. Marcar tarea como completada
    4. Eliminar tarea
    5. Salir
    Opción:
    ```
