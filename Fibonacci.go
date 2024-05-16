package main

import "fmt"

func main() {

	var (
		Terms   int64
		Number1 int64
		Number2 int64
	)

	fmt.Println("We are going to Print the Fibonacci Series!!!\n")

	fmt.Print("Enter the Number of Terms in the Series : ")
	fmt.Scanln(&Terms)
	fmt.Println()

	fmt.Print("Enter the First Term : ")
	fmt.Scanln(&Number1)

	fmt.Print("Enter the Second Term : ")
	fmt.Scanln(&Number2)
	fmt.Println()

	fmt.Printf("The Fibonacci Series of %v Terms starting with %v and %v are : \n\n", Terms, Number1, Number2)

	Fibonacci(Terms, Number1, Number2)

}

func Fibonacci(Terms int64, Num1 int64, Num2 int64) {

	for x := int64(0); x < Terms; x++ {

		fmt.Println(Num1)
		Num1, Num2 = Num2, Num1+Num2

	}

}
