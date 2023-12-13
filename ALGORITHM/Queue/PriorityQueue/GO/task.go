package main

type Task struct {
	tasks    chan func()
	priority uint8
	extra    string
}

func (t *Task) Push(method func()) {
	t.tasks <- method
}

func (t *Task) Get() func() {
	return <-t.tasks
}

func InitTask(priority uint8) *Task {
	return &Task{
		tasks:    make(chan func()),
		priority: priority,
	}
}
