package main

import "fmt"

var Data = make(map[string]string)

func main() {

	var (
		Terms int
	)

	fmt.Println("We are going to store the Name of the Teacher and the Subject!!!\n")

	fmt.Print("Enter the Number of Records : ")
	fmt.Scanln(&Terms)
	fmt.Println()

	Record(Terms)
	Print()

}

func Record(Terms int) {

	var (
		Name    string
		Subject string
	)

	for x := 0; x < Terms; x++ {

		fmt.Print("Enter the Name of the Teacher (Without Spaces use Underscore instead) : ")
		fmt.Scanln(&Name)

		fmt.Print(Name + " teaches the Subject (Without Spaces use Underscore instead) : ")
		fmt.Scanln(&Subject)

		Data[Name] = Subject

	}

	fmt.Println()

}

func Print() {

	fmt.Println("Printing the Stored Records : \n")

	for Key, Value := range Data {

		fmt.Printf("%v teaches %v\n", Key, Value)

	}

}
