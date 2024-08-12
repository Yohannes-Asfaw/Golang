package data

import (
    "errors"
    "task4/models"
    "time"
	"fmt"
)

var tasks = make(map[string]models.Task)

func GetTasks() []models.Task {
    taskList := []models.Task{}
    for _, task := range tasks {
        taskList = append(taskList, task)
    }
    return taskList
}

func GetTaskByID(id string) (models.Task, error) {
    task, exists := tasks[id]
    if !exists {
        return models.Task{}, errors.New("task not found")
    }
    return task, nil
}

func CreateTask(task models.Task) {
    task.ID = generateID() 
    task.DueDate = time.Now().AddDate(0, 0, 7) 
    tasks[task.ID] = task
}

func UpdateTask(id string, updatedTask models.Task) error {
    _, err := GetTaskByID(id)
    if err != nil {
        return err
    }
    updatedTask.ID = id
    tasks[id] = updatedTask
    return nil
}

func DeleteTask(id string) error {
    _, err := GetTaskByID(id)
    if err != nil {
        return err
    }
    delete(tasks, id)
    return nil
}

func generateID() string {
    return fmt.Sprintf("%d", time.Now().UnixNano())
}
