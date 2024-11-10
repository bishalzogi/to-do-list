package main

import "fmt"

func main() {
	printItems()
}

func printItems() {
	var firstTask = "learn go"
	var secondTask = "Do projects in go"
	var thrdTask = "learn Azure"

	var taskItems = []string{firstTask, secondTask, thrdTask}
	fmt.Println("This is the list of Today's TO-DO-List Task")

	for index, task := range taskItems {

		fmt.Printf("%d. %s\n", index+1, task)
	}
}
