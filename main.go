// TASK
// created<DateTime>
// due<DateTime>
package main

import "time"

// let's store these as strings tho
type task struct {
	name    string
	created time.Time
	due     time.Time
}

func newTask(due string, name string) (*task, error) {
	pDue, err := time.Parse(time.RFC3339, due)
	if err != nil {
		return nil, err
	}

	t := task{
		due:  pDue,
		name: name,
	}

	t.created = time.Now()

	return &t, nil
}
