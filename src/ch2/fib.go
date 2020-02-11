// Computing the n-th Fibonacci number iteratively.
package main

import "fmt"

func main() {
	var n int = 10
	fmt.Println(fib(n))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
