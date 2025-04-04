package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	//* use an anonymous function
	transformed := transformNumbers(&numbers, func(val int) int {
		return val * 2
	})
	fmt.Println(transformed)

	//* use a closure
	doubleFn := createTransformer(2)
	tripleFn := createTransformer(3)
	doubledNumbers := transformNumbers(&numbers, doubleFn)
	tripledNumbers := transformNumbers(&numbers, tripleFn)
	fmt.Println("doubled and tripled numbers using Closures")
	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// * Closure
func createTransformer(factor int) func(int) int {
	return func(val int) int {
		return val * factor
	}

}
