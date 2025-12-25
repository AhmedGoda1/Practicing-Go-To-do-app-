package main

import "fmt"

func main() {
	fmt.Println("Hello! Welcome to my ToDo app!")
	fmt.Println("Here's a list of my ToDo: \n")

	shortGolang := "Watch Go crash course."
	fullGolang := "Watch Nana's full course."
	rewardDessert := "Reward myself with a donut."
	listOfTasks := []string{shortGolang, fullGolang, rewardDessert}

	printTasks(listOfTasks)
	fmt.Println("\n")
	listOfTasks = addTask(listOfTasks, "Nero7 nel3ab beng")
	printTasks(listOfTasks)

	printTasks(listOfTasks)
}

func printTasks(taskItems []string) {
	for index, tasks := range taskItems {
		fmt.Printf("%d: %s\n", index+1, tasks)
	}
	return
}
func addTask(taskItems []string, newTask string) []string {
	var updatedList = append(taskItems, newTask)
	return updatedList
}
