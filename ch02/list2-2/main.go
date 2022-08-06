package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a price")
	// n, err := strconv.Atoi(x)
	n, _ := strconv.Atoi(x)
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
