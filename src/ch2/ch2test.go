package main

import "fmt"

var global *int

func main() {
	f()
	fmt.Println(*global)
}

func f() {
	var x int
	x = 1
	global = &x
}
