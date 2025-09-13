package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// task represents a single task with a name, description, and completion status.
// This struct is not directly used in the current implementation, which interacts
// with a database, but it's kept for potential future in-memory operations.
type task struct {
	name        string
	description string
	completed   bool
}

// main is the entry point of the application.
// It initializes the database connection and presents a menu to the user
// for interacting with the task list.
func main() {
	// Establish a connection to the MySQL database.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Loop indefinitely to show the menu until the user decides to exit.
	for {
		fmt.Println("Selecciona la opción deseada:")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tarea")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Opción: ")

		// Read user input from the console.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		// Perform an action based on the user's choice.
		switch choice {
		case 1:
			PrintList(db)
		case 2:
			InsertList(db)
		case 3:
			CheckList(db)
		case 4:
			DeleteList(db)
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		}
	}
}

// PrintList retrieves and displays all tasks from the database.
//
// It takes a database connection as a parameter.
// - db: A pointer to the sql.DB object.
func PrintList(db *sql.DB) {
	// Query the database for all tasks.
	rows, err := db.Query("SELECT id, tasks, description, completed FROM task_list")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Print each task to the console.
	fmt.Println("Tareas:")
	for rows.Next() {
		var id int
		var tasks string
		var description string
		var completed bool

		err := rows.Scan(&id, &tasks, &description, &completed)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%d. %s: %s (Completado: %t)\n", id, tasks, description, completed)
	}
}

// InsertList adds a new task to the database.
//
// It prompts the user for the task's name and description,
// and then inserts the new task into the database.
// It takes a database connection as a parameter.
// - db: A pointer to the sql.DB object.
func InsertList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	// Get the task name and description from the user.
	fmt.Print("Nombre de la tarea: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Descripción de la tarea: ")
	scanner.Scan()
	description := scanner.Text()

	// Insert the new task into the database.
	_, err := db.Exec("INSERT INTO task_list (tasks, description, completed) VALUES (?, ?, ?)", name, description, false)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Tarea agregada con éxito.")
}

// CheckList marks a task as completed in the database.
//
// It prompts the user for the ID of the task to mark as completed,
// and then updates the task's status in the database.
// It takes a database connection as a parameter.
// - db: A pointer to the sql.DB object.
func CheckList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	// Get the ID of the task to mark as completed from the user.
	fmt.Print("Introduce el ID de la tarea completada: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	// Update the task's status in the database.
	_, err := db.Exec("UPDATE task_list SET completed = ? WHERE id = ?", true, id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Tarea con ID %d completada\n", id)
}

// DeleteList removes a task from the database.
//
// It prompts the user for the ID of the task to delete,
// and then removes the task from the database.
// It takes a database connection as a parameter.
// - db: A pointer to the sql.DB object.
func DeleteList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	// Get the ID of the task to delete from the user.
	fmt.Print("Introduce el ID de la tarea que deseas eliminar: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	// Delete the task from the database.
	_, err := db.Exec("DELETE FROM task_list WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("La tarea con ID %d ha sido eliminada\n", id)
}
