package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// n - number of shares to generate
// k - number of shares needed to reconstruct the secret
// TODO: Use a more secure random number generator
func makeShares(secret, k, n int) []int {
	var curve = make([]int, k)
	for i := 0; i < k; i++ {
		if i == 0 {
			curve[i] = secret
			continue
		}
		curve[i] = rand.Intn(secret)
	}
	fmt.Println(curve)
	var shares = make([]int, n)
	for i := 0; i < n; i++ {
		shares[i] = evaluatePolynomial(curve, i+1)
	}
	return shares
}

func evaluatePolynomial(polynomial []int, point int) int {
	var result int
	for i := 0; i < len(polynomial); i++ {
		result += polynomial[i] * powInt(point, i)
	}
	return result
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	fmt.Println(makeShares(1234, 2, 10))
}
