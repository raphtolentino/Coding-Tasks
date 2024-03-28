package main

import (
	"fmt"
	"math"
)

// for loop (1)
// Go only have [For] loops
func for1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// for (2)
// int and post are optional

func for2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// for (3)
// [;] can be used for while

func for3() {
	sum := 1
	for sum < 1090 {
		sum += sum
	}
	fmt.Println(sum)
}

// forever (4)
// if the loop condition is removed, it will keep on looking forever

func foreverLoop() {
	for {
	}
}

// if (5)
// [If] in Go are like (for) loops. It cant be surrounded by () but {} are required

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func if1() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// if with a short statement 


