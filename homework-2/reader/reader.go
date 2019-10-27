package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
Loop pending integer input
*/
func EnterInteger(message string) int64 {
	i := 0
	for {
		if i == 5 {
			panic(fmt.Sprintf("If you didn’t have enough %v attempts, then you’re probably a robot!", i))
		}
		num, err := strconv.ParseInt(getData(message), 10, 64)
		if err == nil {
			return num
		}
		fmt.Println("Is not integer number, please try again")
		i++
	}
}

/**
Get data from user input
*/
func getData(message string) (data string) {
	fmt.Print(message)
	_, _ = fmt.Scanln(&data)

	return
}

/**
Delay execution until 'Enter' button is pressed.
*/
func EnterWaiting() {
	fmt.Print("Press 'Enter' to continue...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}
