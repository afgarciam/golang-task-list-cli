package services

import (
	"fmt"
	"log"

	"todo_cli/ado"
)

func Delete(id int) bool {
	db, err := ado.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := `DELETE FROM todo WHERE id = ?`

	result, err := db.Exec(sql, id)

	if err != nil {
		log.Println(err)
		return false
	}

	if rows, _ := result.RowsAffected(); rows > 0 {
		fmt.Println("Task deleted succesfully!")
		return true
	}

	return false
}
