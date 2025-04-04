package main

func main() {

	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	println(sum)

	// * variadic functions
	println(sumupVariadic(2, 4, 18))
	// * passing a slice to a variadic function
	println(sumupVariadic(numbers...))

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
