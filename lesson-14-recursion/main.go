package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(1))
	fmt.Println(factorial(3))

}

// factorial of  5 => 5 * 4 * 3 * 2 * 1 => 120
// factorial of  4 => 4 * 3 * 2 * 1 => 24
// factorial of  3 => 3 * 2 * 1 => 6

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	// We're calling the function inside the function, this is called recursion
	return n * factorial(n-1)

	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }
	// return result
}
