// Fibonacci Closure program

package main

import "fmt"

// fib returns a function that returns a good numbers.

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f(), f())

}
