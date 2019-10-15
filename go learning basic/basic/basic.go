package main

import (
	"fmt"
	"time"
)

//-----------------------------------------------------------------function sample-----------------------------------------------------------------
// func example with return type
func addFuncWithReturnType(a int, b int) int {
	return a + b
}

// func example with return type
func subFuncWithReturnType(a int, b int) int {
	return a - b
}

// func example without return type
func printResultFuncWithoutReturn(a int, b int, r int) {

	fmt.Println("Sum = ", r)
	x := fmt.Sprintf("%d + %d = %d", a, b, r) // string interpolation
	fmt.Println(x)
}

//function pointer
type calcFunc func(int, int) int

func calc(a int, b int, f calcFunc) int {
	r := f(a, b)
	return r
}

// function with multiple return type
func calcWithMultipleReturnType(a int, b int) (int, int) {
	add := a + b
	sub := a - b
	return add, sub
}

// anonymous function
func anonymousSum(a int, b int) {
	sum := func(x int, y int) int {
		return (x * a) + (y * b)
	}(2, 3)
	fmt.Println("Anonymous sum: ", sum)
}

//-----------------------------------------------------------------conditional sample-----------------------------------------------------------------
// if-else example
func oddEvenUsingIfElse(r int) {
	if r%2 == 0 {
		fmt.Printf("%d is even\n", r)
	} else {
		fmt.Printf("%d is odd\n", r)
	}
}

// switch-case example
func dayPrintUsingSwitchCase() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("today is weekend!")
	default:
		fmt.Println("today is weekday!")
	}
}

//-----------------------------------------------------------------loop sample-----------------------------------------------------------------
// for loop example
func sumUsingForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 9 = ", sum)
}

// while loop example
func sumUsingWhileLoop() {
	sum := 0
	for sum < 10 {
		sum++
	}
	fmt.Println("Sum of 1 to 9 = ", sum)
}

// foreach example
func printForEachItem() {
	arrayOne := [3]string{"Apple", "Mango", "Banana"}

	// print index and value
	for index, element := range arrayOne {
		r := fmt.Sprintf("At index = %d, value is = %s", index, element)
		println(r)
	}

	// print index only
	for index := range arrayOne {
		fmt.Println("Index = ", index)
	}

	// print value only
	for _, element := range arrayOne {
		fmt.Println("Element = ", element)
	}
}

//-----------------------------------------------------------------main function-----------------------------------------------------------------
func main() {
	// var a, b int // another form of declaration

	var a = 10 // variable declaration
	b := 50    // Short declaration

	// func declaration sample ------------------------------------------------------>>>>>>>>>>>>>>>>>>>>>>>>>>
	printResultFuncWithoutReturn(a, b, addFuncWithReturnType(a, b))

	res := calc(15, 5, subFuncWithReturnType)
	fmt.Println("function pointer called: ", res)

	add, sub := calcWithMultipleReturnType(100, 50)
	fmt.Println("multiple return values are: ", add, sub)

	anonymousSum(20, 30)

	// conditional sample ----------------------------------------------------------->>>>>>>>>>>>>>>>>>>>>>>>>>
	oddEvenUsingIfElse(addFuncWithReturnType(a, b))
	dayPrintUsingSwitchCase()

	// loop sample ------------------------------------------------------------------>>>>>>>>>>>>>>>>>>>>>>>>>>
	sumUsingForLoop()
	sumUsingWhileLoop()
	printForEachItem()
}
