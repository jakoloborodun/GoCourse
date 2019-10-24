package converter

import "fmt"

const rateUSD float32 = 63.9

/**
Convert RUB to USD by defined rate.
*/
func RubToUSD() {
	var rub float32
	fmt.Println("Convert RUB to USD")
	fmt.Print("Please enter the amount in RUB to convert to USD: ")
	fmt.Scan(&rub)

	usd := rub / rateUSD

	fmt.Printf("Your amount after conversion: $%0.2f USD\n", usd)
}
