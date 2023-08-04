package main

import "fmt"

func main() {

	var (
		Num1 int
		Num2 int
	)

	fmt.Print("Enter the First Number : ")
	fmt.Scanln(&Num1)

	fmt.Print("Enter the Second Number : ")
	fmt.Scanln(&Num2)
	fmt.Println()

	fmt.Println("Number 1 before Swapping : ", Num1)
	fmt.Println("Number 2 before Swapping : ", Num2)

	Num1, Num2 = Num2, Num1

	fmt.Println("Number 1 after Swapping : ", Num1)
	fmt.Println("Number 2 after Swapping : ", Num2)

}
