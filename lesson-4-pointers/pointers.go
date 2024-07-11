package main

import "fmt"

func main() {

	// This is only for tutorial, That simple value not worth to use pointers.
	age := 32 // regular variable

	// using pointer as value name just for tutotial (We may use e.x userAge)
	agePointer := &age

	// To get value of pointer, we should add "*" before its name
	fmt.Println("Age: ", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Age: ", adultYears)

}

// Also we can use pointers in functions
func getAdultYears(age *int) int {
	return *age - 18

	// override
	// *age = *age - 18
}
