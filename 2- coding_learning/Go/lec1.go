package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// print favorite number (1)

func favInt() {
	fmt.Println("My favorite number is",
		rand.Intn(10))
}

// how to calculate pi (3)

func pi() {
	fmt.Println(math.Pi)
}

// how to add two numbers (4)

func add(x int, y int) int {
	return x + y
}

func addInt() {
	fmt.Println(add(42, 13))
}

// how to add two numbers with a short parameters (5)

func add2(x, y int) int {
	return x + y
}
func add2result() {
	fmt.Println(add(42, 13))
}

// multiple results (6)

func swap(x, y string) (string, string) {
	return y, x
}

func swapPrint() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// named return values (7)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func splitPrIn() {
	fmt.Println(split(17))
}

// variables (8)

// In Go, var statement can be at package or a function level

var c, python, java bool

func lanBool() {
	var i int
	fmt.Println(i, c, python, java)
}

// variables with initializers

// var declarations can inclue initializers, one per variable
// if is present, the type can be removed, the variable will take the type of the inilializer

var i, j int = 1, 2 // i and j are int, and they represent 1 and 2

func varIntz() {
	var c, python, java = true, false, "no!" // variables are between all 3
	fmt.Println(i, j, c, python, java)
}

// short variable declarations (8 - 10)
// using the symbol (:=) means the same thing as var with implicit type
// outside a function, (:) is not avaliable

func shVar() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

// basic types (11)
// bool = boolean
// string = string
// int8 - int64 = integer
// unit - uint64 = unisigned int
// byte /= used for uint8
// rune /= used for int32 (unicode code point)
// float32 and float64 = float
// complex64 and comlex128 = complex (compose of other existing other data types)

var (
	ToBe   bool       = false     // variable ToBe is a boolean that is false
	MaxInt uint64     = 1<<64 - 1 // variable MaxInt is a unit64
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func types() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Valye: %v\n", z, z)
}

// zero value (12)
// zero values is used for
// 0 for numeric types
// false for the boolean type
// "" (the empty string) for strings.

func zValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// type conversions (13)
// The expression T(v) convert the value (v) to the type (T)
// numeric convesion:
// var i int = 42
// var f float64 = float64(i)
// var u unit = unit(f)
// made easy
// i := 42
// f:= float64(i)
// u := uint(f)
func tConv() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

}

// type inference (14)
// when declaring a var without specifying its data type [:= or var]
// the variable type is inferred from the value on the right side

func changes() {
	v := 42 //change me
	fmt.Printf("v is of the type %T\n", v)
}

// constants (15)
// [const] variable being char , string , bool, or numberic numeric
// [:=] cannot be used

const PI = 3.14

func nconst() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")

	const Truth = true
	fmt.Println("Go rules ?", "Truth")
}

// numeric constants (16)
// high precision values

const (
	// Create a huge number by shifiting 1 left by 100 places
	Big = 1 << 100
	// shift again 99 to the right, to changed back into the original number
	Small = Big >> 99
)
func needInt(x int) int { return x*10+1}
func needFloat(x float64) float64{
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}



