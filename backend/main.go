package main

import "fmt"

func main() {
	result, _ := add(14, 73)

	fmt.Printf("Hello World %v", result)

}

func add(x int, y int) (int, bool) {
	return x + y, x < y
}
