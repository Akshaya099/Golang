package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	result := add(5, 7)
	fmt.Println("5 + 7 =", result)

	a, b := swap("hello", "world")
	fmt.Println("Swapped values:", a, b)

	sayHello()
}

func sayHello() {
	fmt.Println("Hello from a simple function!")
}