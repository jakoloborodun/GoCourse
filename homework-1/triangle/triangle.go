package triangle

import (
	"fmt"
	"math"
)

/**
Calculate an area, perimeter and hypotenuse for a right triangle.
*/
func Calculate() {
	var legA, legB float64

	fmt.Println("Calculation of the area, perimeter and hypotenuse of a right triangle.")
	fmt.Print("Please enter the length of leg A: ")
	fmt.Scan(&legA)
	fmt.Print("Please enter the length of leg B: ")
	fmt.Scan(&legB)

	hypotenuse := hypotenuse(legA, legB)
	area := area(legA, legB)
	perimeter := perimeter(legA, legB, hypotenuse)

	fmt.Printf("Right triangle with legs a = %0.2f and b = %0.2f \nhas hypotenuse c = %0.2f\nhas area = %0.2f\nhas perimeter = %0.2f\n", legA, legB, hypotenuse, area, perimeter)
}

/**
Calculate triangle area by given legs.
*/
func area(legA float64, legB float64) float64 {
	return legA * legB / 2
}

/**
Calculate triangle perimeter by its sides.
*/
func perimeter(legA float64, legB float64, hypotenuse float64) float64 {
	return legA + legB + hypotenuse
}

/**
Calculate the hypotenuse of the triangle by the given legs.
*/
func hypotenuse(legA float64, legB float64) float64 {
	return math.Sqrt(legA*legA + legB*legB)
}
