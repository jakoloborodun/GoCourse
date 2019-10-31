package fibonacci

import "math/big"

// Fibonacci is a function that returns
// a function that returns an int value.
func Fibonacci() func() *big.Int {
	x := big.NewInt(-1)
	y := big.NewInt(1)

	return func() (ret *big.Int) {
		ret = x
		x.Add(x, y)
		x, y = y, x
		return
	}
}
