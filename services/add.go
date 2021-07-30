package services

import (
	"fmt"
	"log"

	"todo_cli/ado"
)

func Add(task string) bool {
	db, err := ado.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := `INSERT INTO todo (task)
			VALUES(?)`

	result, err := db.Exec(sql, task)

	if err != nil {
		log.Println(err)
		return false
	}

	if rows, _ := result.RowsAffected(); rows > 0 {
		fmt.Println("Task added succesfully!")
		return true
	}

	return false
}
