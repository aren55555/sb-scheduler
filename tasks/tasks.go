package tasks

type Tasks []Task

type Task struct {
	Done   chan interface{}
	action func()
}

func New(f func()) *Task {
	return &Task{
		Done:   make(chan interface{}, 1),
		action: f,
	}
}

func (t *Task) Execute() {
	t.action()
	t.Done <- nil
}
