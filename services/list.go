package services

import (
	"log"
	"os"
	"time"
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
				END complete,
				created_at
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
	var unixEpoch int64

	for rows.Next() {
		t := new(models.Task)
		rows.Scan(&t.Id, &t.Task, &t.Complete, &unixEpoch)

		t.CreatedAt = time.Unix(unixEpoch, unixEpoch)

		tasks = append(tasks, *t)
	}

	printTable(tasks)
}

func printTable(tasks []models.Task) {
	tableOut := table.NewWriter()
	tableOut.SetOutputMirror(os.Stdout)
	tableOut.AppendHeader(table.Row{"ID", "DATE", "TASK", "COMPLETE"})

	for _, e := range tasks {
		if e.Id != 0 {
			tableOut.AppendRow(table.Row{e.Id, e.CreatedAt.Format(time.RFC3339), e.Task, e.Complete})
		}
	}

	tableOut.Render()
}
