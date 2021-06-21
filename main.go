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

func constructSecret(shares []Point, prime int) int {
	fmt.Println("Prime", prime)
	xs, ys := extractCordinates(shares)
	x := 0
	result := 0
	for i := 0; i < len(ys); i++ {
		currProduct := 1
		for j := 0; j < len(xs); j++ {
			if i != j {
				ai := (x - xs[j])
				bi := (xs[i] - xs[j])
				fmt.Println("a,b", ai, bi)
				a := mod(ai, prime)
				b := mod(bi, prime)
				fmt.Println("mod a,b", a, b)
				_, _, bInverse := extendedGcd(prime, b)
				bInverse = mod(bInverse, prime)
				fmt.Println(b, bInverse)
				c := mod(a*bInverse, prime)
				fmt.Println(c)
				currProduct *= c
			}
		}
		result += ys[i] * currProduct
	}
	return mod(result, prime)
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

// This will be used to compute inverse of an element a
// in a finite field n
// t = inverse of a in field n
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

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	shares := makeShares(35, 2, 10)
	fmt.Println("Shares", shares)
	secret := constructSecret(shares, prime)
	fmt.Println(secret)
}
