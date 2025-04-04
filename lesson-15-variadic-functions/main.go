package main

import "fmt"

func main() {

	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	fmt.Println(sum)

	// * variadic functions
	fmt.Println(sumupVariadic(2, 4, 18))
	// * passing a slice to a variadic function
	fmt.Println(sumupVariadic(numbers...))

}

func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// ! creating variadic functions using ellipsis
func sumupVariadic(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
