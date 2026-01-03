package main

import (
	"fmt"
	"tasks/tasks_components"
)

func main() {
	taskList, err := tasks_components.ReadTasksJSON()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		var option uint8
		fmt.Println("===== SUPER TASKS =====")
		fmt.Println("1. Add new task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Remove tasks")
		fmt.Println("0. Exit")
		fmt.Print("Option> ")
		fmt.Scan(&option)
		switch option {
		case 1:
			taskList, err = tasks_components.AddTaskUI(taskList)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		case 2:
			err = tasks_components.PrintTaskListUI(taskList)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		case 3:
			taskList, err = tasks_components.DeleteTaskUI(taskList)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if option == 0 {
			fmt.Println("Goodbye!")
			break
		}
	}
}
