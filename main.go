package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: Use a more secure random number generator
func makeShares(secret int, minimum int, shares int) []int {
	var polynomial = make([]int, minimum+1)
	polynomial[0] = secret
	for i := 0; i < minimum; i++ {
		polynomial[i+1] = rand.Intn(secret)
	}
	fmt.Println(polynomial)
	return []int{}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	fmt.Println(makeShares(1234, 2, 10))
}
