package models

type Runner struct {
	tasks  []*Task
	params map[string]interface{}
}

func (r *Runner) Invoke(buildFile string, tasks []string, parameters map[string]interface{}) {
	if len(tasks) > 0 {
		for task := range tasks {
			InvokeTask(task)
		}
	} else {
		InvokeTask("default")
	}

}

func (runner *Runner) InvokeTask(name string) {
	currentTask := runner.getTask(name)
	if currentTask == nil {
		//
	}

	if len(runner.dependsOn) > 0 {
		for depend := range runner.dependsOn {
			runner.InvokeTask(depend.name)
		}
	}

	if currentTask.preaction != nil {
		currentTask.preaction()
	}

	currentTask.action()

	if currentTask.postaction != nil {
		currentTask.postaction()
	}
}

func (runner *Runner) getTask(name string) *Task {
	for task := range runner.tasks {
		if task.name == name {
			return task
		}
	}
	return nil
}
