package main

import (
	"fmt"
	"time"

	"github.com/aren55555/superblocks/scheduler"
)

func main() {
	fmt.Println("program started")

	s := scheduler.New()

	s.Add(func() {
		fmt.Println("first task")
	}, 5*time.Second)

	s.Add(func() {
		fmt.Println("second task")
	}, 2*time.Second)

	s.Do()
}
