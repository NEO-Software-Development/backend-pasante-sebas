package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type task struct {
	name        string
	description string
	completed   bool
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for {
		fmt.Println("Selecciona la opción deseada:")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tarea")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Opción: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

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

func PrintList(db *sql.DB) {
	rows, err := db.Query("SELECT id, tasks, description, completed FROM task_list")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

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

func InsertList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nombre de la tarea: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Descripción de la tarea: ")
	scanner.Scan()
	description := scanner.Text()

	_, err := db.Exec("INSERT INTO task_list (tasks, description, completed) VALUES (?, ?, ?)", name, description, false)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Tarea agregada con éxito.")
}

func CheckList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Introduce el ID de la tarea completada: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	_, err := db.Exec("UPDATE task_list SET completed = ? WHERE id = ?", true, id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Tarea con ID %d completada\n", id)
}

func DeleteList(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Introduce el ID de la tarea que deseas eliminar: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	_, err := db.Exec("DELETE FROM task_list WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("La tarea con ID %d ha sido eliminada\n", id)
}
