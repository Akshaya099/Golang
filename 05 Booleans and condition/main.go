package main

import "fmt"

func main() {
	// intergers

	marks := 78

	fmt.Println("Marks:", marks)
	fmt.Println("Is marks >= 90?", marks >= 90)
	fmt.Println("Is marks < 50?", marks < 50)
	fmt.Println("Is marks != 75?", marks != 75)

	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Strings
	departments := []string{"CSE", "IT", "AIML", "EEE", "ECE"}

	for index, dept := range departments {
		if index == 1 {
			fmt.Println("Skipping department at position", index, "-", dept)
			continue
		}
		if index > 3 {
			fmt.Println("Stopping at department position", index, "-", dept)
			break
		}
		fmt.Println("Department at position", index, "-", dept)
	}

}
