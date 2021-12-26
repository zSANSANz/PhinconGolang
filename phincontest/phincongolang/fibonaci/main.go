package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	var a int = 0

	var b int = 1

	temp := 0
	for i := 2; i < n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return b
}
