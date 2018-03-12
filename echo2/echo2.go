// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	secs := time.Since(start)
	fmt.Println(s, secs)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
