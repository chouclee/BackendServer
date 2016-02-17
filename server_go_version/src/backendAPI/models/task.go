package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Tasks map[string]*Task
)

type Task struct {
	TaskId    string
	UserId    string
	ModelTpye string
}

func init() {
	Tasks = make(map[string]*Task)
}

func AddTask(task Task) (TaskId string) {
	task.TaskId = "honeycomb" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Tasks[task.TaskId] = &task
	return task.TaskId
}

func GetTask(TaskId string) (task *Task, err error) {
	if v, ok := Tasks[TaskId]; ok {
		return v, nil
	}
	return nil, errors.New("TaskId Not Exist")
}

func GetAllTasks() map[string]*Task {
	return Tasks
}

func DeleteTask(TaskId string) {
	delete(Tasks, TaskId)
}
