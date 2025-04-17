package main

import "fmt"
func main(){
	// define variables 
	var first_name string
	var last_name string
	var age int
	var weight float64

	//set value for each variable:
	first_name = "John"
	last_name = "wick"
	age = 44
	weight = 80.50

	// use variables
	fmt.Println("first name is:", first_name, ".")
	fmt.Println("last name is:", last_name, ".")
	fmt.Println(first_name, "is", age, "years old.")
	fmt.Println(first_name, "is", weight, "KG.")
	fmt.Println("----END----")
}
