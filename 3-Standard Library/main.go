package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Lets start working"

	fmt.Println(strings.ToUpper(text)) 
	fmt.Println(strings.ToLower(text)) 
	fmt.Println(strings.ReplaceAll(text, "working", "learning")) 
	fmt.Println(strings.Contains(text, "to")) 
	fmt.Println(text) 

	nums := []int{10, 50, 30, 20, 40}
	sort.Ints(nums)
	fmt.Println("Sorted Array:", nums)

	index :=sort.SearchInts (nums, 30)
	fmt.Println(index)

}
