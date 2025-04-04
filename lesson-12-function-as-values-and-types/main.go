package main

import "fmt"

type transformFn func(int) int // create a new type for a function that takes an int and returns an int

func main() {

	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	// get the function as a value
	transformFn1 := getTransformFn(&numbers)
	transformFn2 := getTransformFn(&moreNumbers)

	// use the function as a value
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformFn2)
	transformNumbers := transformNumbers(&numbers, transformFn1)

	fmt.Println(transformNumbers)
	fmt.Println(transformedMoreNumbers)

}

// untility function to transform the numbers, We're using a function as a parameter
func transformNumbers(n *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *n {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// utility function to return a function as a value
func getTransformFn(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// utility function to double the numbers
func double(number int) int {
	return number * 2
}

// utility function to triple the numbers
func triple(number int) int {
	return number * 3
}
