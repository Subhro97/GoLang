package main

import (
	"fmt"

	"github.com/Arrays/product"
)

func main() {
	hobbies := [3]string{"Play Video Games", "Go for brisk walk", "Painting Shree Ram"}

	fmt.Println("Hobbies: ", hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	slicedHobbies := hobbies[0:2]
	slicedHobbies2 := hobbies[:2]
	fmt.Println("Sliced Hobbies: ", slicedHobbies)
	fmt.Println("Sliced Hobbies 2: ", slicedHobbies2)

	slicedHobbies = slicedHobbies[1:3]
	fmt.Println("Re-sliced Hobbies: ", slicedHobbies)

	goals := []string{"Be System Architect", "Will be re-assigned"}

	fmt.Println("Goals: ", goals)

	goals[1] = "Have good Investment Skills"

	updatedGoals := append(goals, "Stay in Top Shape - Fitness, Strength & Attractive")

	fmt.Println(updatedGoals)

	tshirt := product.New("T-Shirt", 1000)
	laptop := product.New("Macbook", 100000)

	productsArr := []product.Product{tshirt, laptop}

	fmt.Println("Products: ", productsArr)

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
//
