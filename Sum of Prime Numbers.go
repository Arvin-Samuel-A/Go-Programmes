package main

import (
	"fmt"
	"math"
	"sync"
)

var Sum uint64 = 2
var Mutex sync.Mutex

func main() {

	var Terms uint64

	fmt.Println("We are going to Calculate the Sum of Prime Numbers till a user-entered limit!!!\n")

	fmt.Print("Enter the Limit : ")
	fmt.Scanln(&Terms)
	fmt.Println()

	Threading(Terms)

	fmt.Printf("The Sum of Prime Numbers till %v is : %v", Terms, Sum)

}

func Threading(Terms uint64) {

	var (
		WG  sync.WaitGroup
		Var float64 = float64(Terms - 2)
	)

	WG.Add(12)
	Var /= 12

	for x := float64(0); x < 12; x++ {

		go Prime_Numbers(uint64(x*Var+3), uint64((x+1)*Var+3), &WG)

	}

	WG.Wait()

}

func Prime_Numbers(Num1 uint64, Num2 uint64, WG *sync.WaitGroup) {

	defer WG.Done()

	var Total uint64

	for x := Num1; x < Num2; x++ {

		if x%2 != 0 {

			var Found bool = true
			var Dup uint64 = uint64(math.Sqrt(float64(x)))

			for y := uint64(3); y <= Dup; y += 2 {

				if x%y == 0 {

					Found = false
					break

				}

			}

			if Found {

				Total += x

			}

		}

	}

	Mutex.Lock()
	Sum += Total
	Mutex.Unlock()

}
