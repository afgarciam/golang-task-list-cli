package services

import (
	"fmt"
	"log"
	"todo_cli/ado"
)

func Complete(id int) bool {
	db, err := ado.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := `UPDATE todo 
			SET complete = 1
			WHERE id = ?`

	result, err := db.Exec(sql, id)

	if err != nil {
		log.Println(err)
		return false
	}

	if rows, _ := result.RowsAffected(); rows > 0 {
		fmt.Printf("Task [%d] completed succesfully! \n", id)
		return true
	}

	return false

}
