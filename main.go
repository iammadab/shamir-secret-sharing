package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// TODO: Use a more secure random number generator
func makeShares(secret int, minimum int, share int) []int {
	var polynomial = make([]int, minimum+1)
	polynomial[0] = secret
	for i := 0; i < minimum; i++ {
		polynomial[i+1] = rand.Intn(secret)
	}
	// Perform validation for number of share
	var shares = make([]int, share)
	for i := 0; i < share; i++ {
		// Evaluates the polynomial f at f(1), f(2), ... f(n)
		// n being the number of shares to be generated
		shares[i] = evaluatePolynomial(polynomial, i+1)
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
