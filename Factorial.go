package main

import "fmt"

func main() {

	var Number int64

	fmt.Println("We are Going to Calculate the Factorial of a user-entered Number!!!\n")

	fmt.Print("Enter the Number to Calculate it's Factorial : ")
	fmt.Scanln(&Number)
	fmt.Println()

	fmt.Println("The Factorial of ", Number, " is : ", Factorial(Number))

}

func Factorial(Num int64) (Output int64) {

	if Num > 1 {

		Output = Num * Factorial(Num-1)

	} else {

		Output = 1

	}

	return
}
