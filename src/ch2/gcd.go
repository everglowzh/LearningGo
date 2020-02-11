// Computing the greatest common divisor (GCD) of two integer.
package main

import "fmt"

func main() {
	var x int = 377
	var y int = 319
	fmt.Println(gcd(x, y))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
