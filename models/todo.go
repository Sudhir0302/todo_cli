package models

import "fmt"

type Todo struct {
	Id     int
	Task   string
	Status bool
}

func (t Todo) String() string {
	return fmt.Sprintf("Task #%d %s %t", t.Id, t.Task, t.Status)
}
