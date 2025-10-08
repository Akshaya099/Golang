package main

import "fmt"

var score = 95.5
var departments = []string{"CSE", "IT", "AIML", "EEE", "ECE"}

func main() {
	sayHello("CSE")

	for _, dept := range departments {
		fmt.Println(dept)
	}

	showScore()
}

func sayHello(dept string) {
	fmt.Println("Welcome to the", dept, "department")
}

func showScore() {
	fmt.Println("Score:", score)
}
