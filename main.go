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

const prime = 7919
const INTERCEPT = 0

// n - number of shares to generate
// k - number of shares needed to reconstruct the secret
// TODO: Use a more secure random number generator
func generateShares(secret, k, n int) []Point {
	var curve = constructPolynomialOfDegree(k)
	curve[INTERCEPT] = secret
	fmt.Println("Curve", curve)
	var shares = pickNPointsFromPolynomial(curve, n)
	return shares
}

func constructPolynomialOfDegree(degree int) []int {
	var polynomial = make([]int, degree)
	for i := 0; i < degree; i++ {
		polynomial[i] = rand.Intn(prime)
	}
	return polynomial
}

func pickNPointsFromPolynomial(polynomial []int, n int) []Point {
	var shares = make([]Point, n)
	for i := 0; i < n; i++ {
		shares[i] = evaluatePolynomial(polynomial, i+1)
	}
	return shares
}

func evaluatePolynomial(polynomial []int, point int) Point {
	var result int
	for i := 0; i < len(polynomial); i++ {
		result += polynomial[i] * pow(point, i)
		result %= prime
	}
	return Point{X: point, Y: result}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func constructSecret(shares []Point) int {
	xs, ys := extractCordinates(shares)
	result := 0
	for i := 0; i < len(ys); i++ {
		currProduct := 1
		for j := 0; j < len(xs); j++ {
			if i != j {
				a := xs[j]
				b := mod((xs[j] - xs[i]))
				c := divmod(a, b)
        currProduct = mod(currProduct * c)
			}
		}
		result += mod(ys[i] * currProduct)
    result = mod(result)
	}
	return mod(result)
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

func divmod(a, b int) int {
	_, _, bInverse := extendedGcd(prime, b)
	bInverse = mod(bInverse)
	return mod(a*bInverse)
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

func mod(a int) int {
	return (a%prime + prime) % prime
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's get started")
	shares := generateShares(32, 2, 10)
	fmt.Println("Shares", shares)
	secret := constructSecret(shares)
	fmt.Println(secret)
}
