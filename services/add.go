package services

import (
	"errors"
	"fmt"
	"log"

	"todo_cli/ado"
)

func Add(task string) error {

	if task == "" {
		return errors.New("don't add a empty task")
	}

	db, err := ado.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := `INSERT INTO todo (created_at,task)
			VALUES(strftime('%s','now'),?)`

	result, err := db.Exec(sql, task)

	if err != nil {
		log.Println(err)
		return err
	}

	if rows, err := result.RowsAffected(); rows > 0 {
		fmt.Println("Task added succesfully!")
		return nil
	} else {
		return err
	}
}
