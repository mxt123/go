package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

//https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func stuff(n int64, k int64) *big.Int {
	defer timeTrack(time.Now(), "stuff")
	r := new(big.Int)
	return r.Binomial(n, k)
}

func main() {
	fmt.Println(stuff(1000, 10))
}
