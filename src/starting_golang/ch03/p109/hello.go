package main

import "fmt"

const ONE = 1

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func main() {
	x, y := one()
	fmt.Printf("x=%d, y=%d\n", x, y)
}
