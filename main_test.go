package main

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	// create a task
	tsk, err := newTask(time.Now().Add(time.Hour).Format(time.RFC3339), "rose")
	if err != nil {
		t.Fatalf("Task was no created")
		return
	}

	// check if task was created
	if tsk.name != "rose" {
		// fail if not
		t.Fatalf("Name is not valid")
	}
}
