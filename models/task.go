package models

import "time"

type Task struct {
	Id        int
	Task      string
	Complete  bool
	CreatedAt time.Time
}
