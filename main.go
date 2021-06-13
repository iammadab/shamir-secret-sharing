package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
  INTERCEPT = 0
)

// TODO: Use a more secure random number generator
func makeShares(secret, minimum, share int) []int {
  var curve = make([]int, minimum)
  for i := 0; i < minimum; i++ {
    if i == 0 {
      curve[i] = secret
      continue
    }
    curve[i] = rand.Intn(secret)
  }
  fmt.Println(curve)
  var shares = make([]int, share)
  for i := 0; i < share; i++ {
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
