package main

import (
	"fmt"
	"time"

	"github.com/aren55555/superblocks/scheduler"
)

type Scheduler interface {
	Schedule(t scheduler.Task, delayDuration time.Duration)
}

var s Scheduler = scheduler.New()

var stop = make(chan int)

func main() {
	fmt.Println("hello world")

	s.Schedule(func() {
		fmt.Println("first task")
	}, 5*time.Second)

	s.Schedule(func() {
		fmt.Println("second task")
	}, 2*time.Second)

	<-stop
}
