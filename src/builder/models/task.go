package models

type Task struct {
	name string
	action func
	preaction func
	postaction func
	dependsOn []Task
}

func NewTask *Task(name string){
	return &Task{name: name}
}

func NewTask(name string) *Task {
	task := new(Task)
	return task
}