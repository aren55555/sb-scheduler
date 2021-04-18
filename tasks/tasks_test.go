package tasks

import (
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	array := []time.Time{}

	testTask := New(func() {
		array = append(array, time.Now())
	})

	go testTask.Execute()
	<-testTask.Done // block until the task has completed

	if got, want := len(array), 1; got != want {
		t.Fatalf("len of results: got %v, want %v", got, want)
	}
}
