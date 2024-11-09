package main

import "fmt"

func main() {
	fmt.Println("To Do List")

	firstTask := "Learn go "
	secondTask := "Learn Azure"
	thirdTask := "Get certified in Azure"

	taskItem := []string{firstTask, secondTask, thirdTask}

	for index, task := range taskItem {

		fmt.Printf("%d: %s\n", index+1, task)
	}
}
