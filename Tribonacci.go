package main

import "fmt"

func main() {

	var (
		Terms   int64
		Number1 int64
		Number2 int64
		Number3 int64
	)

	fmt.Println("We are going to print the Tribonacci Series!!!\n")

	fmt.Print("Enter the Number of Terms : ")
	fmt.Scanln(&Terms)
	fmt.Println()

	fmt.Print("Enter the First Number : ")
	fmt.Scanln(&Number1)

	fmt.Print("Enter the Second Number : ")
	fmt.Scanln(&Number2)

	fmt.Print("Enter the Third Number : ")
	fmt.Scanln(&Number3)
	fmt.Println()

	fmt.Printf("The Tribonacci Series for %v terms with First term %v, Second term %v and Third term %v is :\n\n", Terms, Number1, Number2, Number3)

	Tribonacci(Terms, Number1, Number2, Number3)

}

func Tribonacci(Terms int64, Num1 int64, Num2 int64, Num3 int64) {

	for x := int64(0); x < Terms; x++ {

		fmt.Println(Num1)

		Num1, Num2, Num3 = Num2, Num3, Num1+Num2+Num3

	}

}
