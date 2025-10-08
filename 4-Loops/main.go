package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fruits := []string{"apple", "banana", "cherry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, value := range fruits {
		fmt.Printf("The fruit at index %v is %v \n", index, value)
	}
}
