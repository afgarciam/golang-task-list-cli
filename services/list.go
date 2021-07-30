package services

import (
	"log"
	"os"
	"todo_cli/ado"
	"todo_cli/models"

	"github.com/jedib0t/go-pretty/table"
)

func List(listAll bool) {
	db, err := ado.GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := `SELECT id, 
				task, 
				CASE WHEN complete = 1 THEN
				 	true 
				ELSE 
					false 
				END complete 
			FROM todo `

	if !listAll {
		sql += " WHERE complete = 0"
	}

	rows, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	tasks := make([]models.Task, 10)

	for rows.Next() {
		t := new(models.Task)
		rows.Scan(&t.Id, &t.Task, &t.Complete)
		tasks = append(tasks, *t)
	}

	tableOut := table.NewWriter()
	// tableOut.SetStyle(table.StyleColoredDark)
	tableOut.SetOutputMirror(os.Stdout)
	tableOut.AppendHeader(table.Row{"ID", "TASK", "COMPLETE"})

	for _, element := range tasks {
		if element.Id != 0 {
			tableOut.AppendRow(table.Row{element.Id, element.Task, element.Complete})

		}
	}

	tableOut.Render()
}
