package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// Task (1)
	fmt.Println("Task 1 Result:")
	hobbies := [3]string{"reading", "coding", "gaming"}
	fmt.Println(hobbies)
	// Task (2)
	fmt.Println("Task 2 Result:")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// Task (3)
	fmt.Println("Task 3 Result:")
	firstHobbies := hobbies[:2]
	lastHobbies := hobbies[1:3]
	fmt.Println(firstHobbies, lastHobbies)

	// Task (4)
	fmt.Println("Task 4 Result:")
	firstHobbies = firstHobbies[1:3]
	fmt.Println(firstHobbies)

	// Task (5)
	fmt.Println("Task 5 Result:")
	courseGoals := []string{"learn Go", "build a web app"}
	fmt.Println(courseGoals)

	// Task (6)
	fmt.Println("Task 6 Result:")
	courseGoals[1] = "Master Golang"
	courseGoals = append(courseGoals, "create WebApp")
	fmt.Println(courseGoals)

	// Task (7)
	fmt.Println("Task 7 Result:")
	products := []Product{
		{"Book", 1, 10.99},
		{"Game", 2, 19.99},
	}
	fmt.Println(products)
	products = append(products, Product{"Music", 3, 29.99})
	fmt.Println(products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
