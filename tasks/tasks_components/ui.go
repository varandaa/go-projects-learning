package tasks_components

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddTaskUI(taskList []Task) ([]Task, error) {
	var err error
	var priority uint8
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Task Description: ")
	desc, err := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)
	if err != nil {
		return nil, err
	}
	fmt.Println()
	fmt.Print("Task Priority: ")
	fmt.Scan(&priority)
	fmt.Println()
	taskList, err = registerTask(desc, priority, taskList)
	if err != nil {
		return nil, err
	}
	//printTask(newTask)
	fmt.Println("Task added successfully!!")
	return taskList, err

}

func printTask(taskToPrint Task) {
	fmt.Printf("Task ID: %v\n", taskToPrint.Id)
	fmt.Printf("Task Description: %s\n", taskToPrint.Description)
	fmt.Printf("Task Priority: %v\n", taskToPrint.Priority)
}

func PrintTaskListUI(taskList []Task) error {
	var err error
	fmt.Println("========================")
	for _, t := range taskList {
		printTask(t)
		if err != nil {
			return err
		}
		fmt.Println("========================")
	}
	return err
}

func DeleteTaskUI(taskList []Task) ([]Task, error) {
	var id int
	var err error
	PrintTaskListUI(taskList)
	fmt.Print("Task ID to delete: ")
	fmt.Scan(&id)
	fmt.Println()
	taskList, err = removeTask(id, taskList)
	if err != nil {
		return nil, err
	}
	return taskList, err
}
