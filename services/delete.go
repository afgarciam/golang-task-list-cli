package services

import (
	"fmt"

	"todo_cli/ado"
)

func Delete(id int) error {
	db, err := ado.GetConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	sql := `DELETE FROM todo WHERE id = ?`

	result, err := db.Exec(sql, id)

	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows > 0 {
		fmt.Println("Task deleted succesfully!")
		return nil
	} else {
		return err
	}

}
