package scheduler

import (
	"fmt"
	"time"
)

type tasks []task

func (t tasks) String() string {
	fmt.Sprintf("")
}

type Task func()

type task struct {
	block    chan interface{}
	RunsAt   time.Time
	Task     Task
	finished chan interface{}
}

type Scheduler struct {
	stop  chan interface{}
	tasks []task
}

func New() *Scheduler {
	return &Scheduler{
		tasks: []task{},
	}
}

func (s *Scheduler) Schedule(t Task, delayDuration time.Duration) {
	newTask := task{
		RunsAt: time.Now().Add(delayDuration),
		Task:   t,
	}

	if len(s.tasks) == 0 {
		s.tasks = append(s.tasks, newTask)
		fmt.Printf("%#v\n", s)
		return
	}

	for idx, t := range s.tasks {
		if t.RunsAt.Before(newTask.RunsAt) {
			continue
		}
		s.tasks = insert(s.tasks, idx+1, newTask)
	}

	fmt.Printf("%#v\n", s)
}

func (s *Scheduler) Do() {
	for {
		select {
		case <-ticker.C:
			go c.fetchAndUpdate()
		case <-c.stop:
			ticker.Stop()
			return
		}
	}
}

func insert(slice []task, index int, value task) []task {
	result := []task{}
	result = append(result, slice[:index+1]...)
	result = append(result, value)
	result = append(result, slice[index:]...)
	return result
}
