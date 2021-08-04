package services

import (
	"fmt"
	"todo_cli/ado"
)

func Complete(id int) error {
	db, err := ado.GetConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	sql := `UPDATE todo 
			SET complete = 1
			WHERE id = ?`

	result, err := db.Exec(sql, id)

	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows > 0 {
		fmt.Printf("Task [%d] completed succesfully! \n", id)
		return nil
	} else {
		return err
	}
}
