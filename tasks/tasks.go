package tasks

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type task struct {
	id     string
	runsAt time.Time
	t      *Task
}

type Tasks struct {
	Next       chan *Task
	ticker     *time.Ticker // used for the first task or if there's space between tasks
	data       []task
	stopTicker chan interface{}
}

func NewTasks() *Tasks {
	return &Tasks{
		Next: make(chan *Task),
		data: []task{},
	}
}

func (t *Tasks) Add(f func(), d time.Duration) {
	if t.ticker != nil {
		t.ticker.Stop()
	}

	// TODO: figure out where to add based on the time
	t.data = append(t.data, task{
		id:     uuid.New().String(),
		runsAt: time.Now().Add(d),
		t:      NewTask(f),
	})

	t.reEnqueue()
}

func (t *Tasks) reEnqueue() {
	firstTask := t.data[0]
	t.ticker = time.NewTicker(firstTask.runsAt.Sub(time.Now()))
	go func() {
		for {
			select {
			case <-t.ticker.C:
				fmt.Println("task started")
				t.pop()
				t.Next <- firstTask.t
				fmt.Println("waiting for task to finish")
				<-firstTask.t.Done
				fmt.Println("task done")
				return
			}
		}
	}()
}

// Pop will return the first element and remove it from the store.
func (t *Tasks) pop() *Task {
	first := t.data[0]
	return first.t
}
