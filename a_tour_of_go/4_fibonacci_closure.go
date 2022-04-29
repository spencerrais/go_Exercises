package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := 0
	prevPrev := 0
	return func() int {
		fib := prev + prevPrev
		if prevPrev == 0 {
			prev = 1
		}
		prevPrev = prev
		prev = fib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
