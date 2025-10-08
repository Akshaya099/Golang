package main

import "fmt"

func main() {
	var Dept [3]string = [3]string{"AIML","IT","CSE"}

	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers, len(numbers))

	cities := [3]string{"Delhi", "Mumbai", "Chennai"}
	fmt.Println(cities)

	cities[1] = "Bangalore" 
	fmt.Println(cities)

	numbers = append(numbers, 90)
	fmt.Println(numbers)

	fruits := []string{"apple", "banana", "cherry", "grape", "mango"}
	part1 := fruits[1:3] 
	part2 := fruits[3:]
	fmt.Println(fruits)
	fmt.Println(part1)
	fmt.Println(part2)
}
