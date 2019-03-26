package main

import (
	"fmt"

	"even"
)

func getNum(k int, getN func(int) int) int {
	return 1 + getN(k)
}

func main() {
	k := 1

	i := getNum(k, even.Double)
	fmt.Println(i)

	i = getNum(k, even.Quadruple)
	fmt.Println(i)
}
