package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	j := strconv.Itoa(i)
	var r [2]int
	r[0] = 5
	r[1] = 6

	const a = 45
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", r[1], r)
	fmt.Printf("%v, %T", a, a)
}
