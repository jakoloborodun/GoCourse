package main

import (
	"GoCourse/homework-2/fibonacci"
	"GoCourse/homework-2/reader"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "End of execution.")
		}
	}()

	fmt.Println("Function to check whether a number is Even or Odd")
	number := reader.EnterInteger("Enter number: ")
	numberIsOddEven(number)

	reader.EnterWaiting()

	fmt.Println()
	fmt.Println("Function to check a number is multiple of denominator")
	numberIsMultipleOfDenominator(number, 3)

	reader.EnterWaiting()

	fmt.Println()
	fmt.Println("100 first fibonacci numbers starting at 0:")
	f := fibonacci.Fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Printf("%v, ", f())
	}
}

/**
Detect given number is odd or even.
*/
func numberIsOddEven(number int64) {
	if number%2 == 0 {
		fmt.Println(number, "is Even number")
	} else {
		fmt.Println(number, "is Odd number")
	}
}

/**
Shows the remainder of dividing the numerator by the denominator.
*/
func numberIsMultipleOfDenominator(number, denominator int64) {
	if remainder := number % denominator; remainder == 0 {
		fmt.Printf("The number %d is divisible without a remainder by the denominator %d\n", number, denominator)
	} else {
		fmt.Printf("The number %v is divided by the denominator %d with the remainder %v\n", number, denominator, remainder)
	}
}
