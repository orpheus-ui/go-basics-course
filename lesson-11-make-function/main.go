package main

import "fmt"

type floatMap map[string]float64 // create a new type for a map of strings and floats

func (fm floatMap) print() {
	for key, value := range fm {
		fmt.Println(key, value)
	}
}

func main() {

	userNames := make([]string, 0, 5) // make a slice of strings with length 2 and capacity 5

	userNames = append(userNames, "John")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Joe")
	userNames = append(userNames, "Stephen")
	userNames = append(userNames, "Chris")

	fmt.Println(userNames)

	courses := make(floatMap, 5) // make a map of strings and floats with capacity 5

	courses["Go"] = 4.7
	courses["Python"] = 4.5
	courses["Java"] = 4.8
	courses["C++"] = 3.9

	courses.print()

}
