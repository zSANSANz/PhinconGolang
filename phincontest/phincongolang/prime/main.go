package main

import (
	"fmt"
	"math"
)

func printPrimeNumbers(num1, num2 int) {
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}

func main() {
	printPrimeNumbers(2, 20)
}
