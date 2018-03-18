package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Println(gcd(strconv.Atoi(os.Args[0]), strconv.Atoi(os.Args[1])))
}
