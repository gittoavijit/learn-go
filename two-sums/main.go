//  Do The sum of two variables taken as input

package main

import "fmt"

func main() {
	var i, j int
	fmt.Println("we are going to do the sum of two variables")
	fmt.Println("Enter number one: ")
	fmt.Scan(&i)
	fmt.Println("Enter number two: ")
	fmt.Scan(&j)
	sum := i + j
	fmt.Printf("The sum of %d and %d is %d \n", i, j, sum)
}
