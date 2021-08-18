package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	j := strconv.Itoa(i)
	fmt.Printf("%v, %T", j, j)
}
