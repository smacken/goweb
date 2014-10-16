package models

type Task struct {
	name       string
	action     func() bool
	preaction  func() bool
	postaction func() bool
	dependsOn  []*Task
}

func NewTask(name string) *Task {
	return &Task{name: name}
}
