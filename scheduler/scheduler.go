package scheduler

import (
	"fmt"
	"time"

	"github.com/aren55555/superblocks/tasks"
)

const (
	refreshInterval = 10 * time.Second
)

type Scheduler struct {
	stop  chan interface{}
	tasks *tasks.Tasks
	next  chan interface{}
}

func New() *Scheduler {
	return &Scheduler{
		stop:  make(chan interface{}, 1),
		tasks: tasks.NewTasks(),
	}
}

func (s *Scheduler) Stop() {
	fmt.Println("stopping...")
	s.stop <- nil
}

func (s *Scheduler) Do() {
	for {
		select {
		case nextTask := <-s.tasks.Next:
			fmt.Println("next task received")
			go nextTask.Execute()
		case <-s.stop:
			return
		}
	}
}

func (s *Scheduler) Add(f func(), d time.Duration) {
	s.tasks.Add(f, d)
}
