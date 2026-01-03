package tasks_components

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Task struct {
	Id          int
	Description string
	Priority    uint8
}

func getNextId(taskList []Task) int {
	return taskList[len(taskList)-1].Id + 1
}

func registerTask(description string, priority uint8, taskList []Task) (updatedTaskList []Task, err error) {
	nextId := getNextId(taskList)
	var taskToRegister Task = Task{nextId, description, priority}
	taskList = append(taskList, taskToRegister)
	err = updateTaskJSON(taskList)
	if err != nil {
		return nil, err
	}
	return taskList, err
}

func removeTask(idToRemove int, taskList []Task) (updatedTaskList []Task, err error) {
	indexToRemove := -1
	for i, task := range taskList {
		if task.Id == idToRemove {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return taskList, fmt.Errorf("task with ID %d not found", idToRemove)
	}
	taskList = slices.Delete(taskList, indexToRemove, indexToRemove+1)
	err = updateTaskJSON(taskList)
	if err != nil {
		return taskList, err
	}
	return taskList, err
}

func updateTaskJSON(taskList []Task) (err error) {
	// fmt.Println(taskList)
	jsonTasks, err := json.Marshal(taskList)
	// fmt.Println(jsonTasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", jsonTasks, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return err
}

func ReadTasksJSON() (taskList []Task, err error) {
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	var jsonTasks []Task
	err = json.Unmarshal([]byte(content), &jsonTasks)
	if err != nil {
		return nil, err
	}
	// fmt.Print(string(content))
	return jsonTasks, nil
}
