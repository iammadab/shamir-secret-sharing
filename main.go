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

const prime = 97

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
		curve[i] = rand.Intn(prime)
	}
	fmt.Println("Curve", curve)
	var shares = make([]Point, n)
	for i := 0; i < n; i++ {
		shares[i] = evaluatePolynomial(curve, i+1, prime)
	}
	return shares
}

func evaluatePolynomial(polynomial []int, point int, prime int) Point {
	var result int
	for i := 0; i < len(polynomial); i++ {
		result += polynomial[i] * int(math.Pow(float64(point), float64(i)))
		result %= prime
	}
	return Point{X: point, Y: result}
}

func constructSecret(shares []Point, prime int) float64 {
	xs, ys := extractCordinates(shares)
	x := 0
	result := 0.0
	for i := 0; i < len(ys); i++ {
		currProduct := 1.0
		for j := 0; j < len(xs); j++ {
			if i != j {
				a := float64(x - xs[j])
				b := float64(xs[i] - xs[j])
				c := a / b
				currProduct *= c
				//fmt.Println(a, b, c)
				//currProduct *= (x - xs[j])/(xs[i] - xs[j])
			}
		}
		result += float64(ys[i]) * currProduct
	}
	return result
}

func extractCordinates(points []Point) ([]int, []int) {
	x := make([]int, len(points))
	y := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		x[i] = points[i].X
		y[i] = points[i].Y
	}
	return x, y
}

func extendedGcd(a, b int) (r, s, t int) {
	// Make sure a is the bigger of the two
	if a < b {
		a, b = b, a
	}
	sa := [...]int{1, 0}
	ta := [...]int{0, 1}
	for b != 0 {
		q := a / b
		a, b = b, a%b
		sa[0], sa[1] = sa[1], sa[0]-q*sa[1]
		ta[0], ta[1] = ta[1], ta[0]-q*ta[1]
	}
	return a, sa[0], ta[0]
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	shares := makeShares(12, 2, 10)
	fmt.Println("Shares", shares)
	secret := constructSecret(shares, prime)
	fmt.Println(secret)
}
