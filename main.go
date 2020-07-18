package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GCD(a, b int, c ...int) int {
	r := gcd(a, b)

	for i, _ := range c {
		r = gcd(r, c[i])
	}

	return r
}

func main() {
	fmt.Println(GCD(12, 24, 36, 48))
}
