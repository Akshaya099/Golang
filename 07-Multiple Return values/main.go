package main

import "fmt"

func calculate(a, b int) (int, int, int) {
	sum := a + b
	product := a * b
	difference := a - b
	return sum, product, difference
}

func main() {
	sum, product, difference := calculate(10, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Difference:", difference)
}