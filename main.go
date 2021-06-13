package main

import (
	"fmt"
	"math"
	"math/rand"
	//"time"
)

type Point struct {
	X int
	Y int
}

// n - number of shares to generate
// k - number of shares needed to reconstruct the secret
// TODO: Use a more secure random number generator
func makeShares(secret, k, n int) []Point {
	var curve = make([]int, k)
	for i := 0; i < k; i++ {
		if i == 0 {
			curve[i] = secret
			continue
		}
		curve[i] = rand.Intn(secret)
	}
	fmt.Println(curve)
	var shares = make([]Point, n)
	for i := 0; i < n; i++ {
		shares[i] = evaluatePolynomial(curve, i+1)
	}
	return shares
}

func evaluatePolynomial(polynomial []int, point int) Point {
	var result int
	for i := 0; i < len(polynomial); i++ {
		result += polynomial[i] * int(math.Pow(float64(point), float64(i)))
	}
	return Point{X: point, Y: result}
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	fmt.Println(makeShares(1234, 2, 10))
}
