package main

import (
	"fmt"
)

//-----------------------------------------------------------------array sample-----------------------------------------------------------------
func arraySample() {
	//var ary [5]int
	ary := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Print array items: ")

	for i, item := range ary {
		r := fmt.Sprintf("At index = %d, value = %d", i, item)
		println(r)
	}
}

//-----------------------------------------------------------------slice sample (list type)------------------------------------------------------
func sliceOperations() {

	integerSlice := []int{10, 20, 30, 40, 50}
	stringSlice := []string{"first", "second", "third", "fourth"}
	fmt.Println("This is the integer slice: ", integerSlice)
	fmt.Println("This is the string slice: ", stringSlice)

	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40

	s = append(s, 60, 70)
	fmt.Println("Added two elements to slice: ", s)

	s = append(s[:2], s[2+1:]...) //... is a variadic argument in Go
	fmt.Println("Deleted one element from slice: ", s)

	s = s[1:4]
	fmt.Println("Slice with second to fourth element: ", s)

	//Get the length of the current slice:
	fmt.Println("Length: ", len(s))

	//Get the capacity of the current slice:
	fmt.Println("Capacity: ", cap(s)) //this is give "7"

	//Copy one slice to another:
	//Make a slice with the same length as "s":
	d := make([]int, len(s))
	copy(d, s)
	fmt.Println("This is the new slice: ", d)
}

func nestedSlice() {
	//Example 1:
	nested := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			out = append(out, j)
		}
		nested = append(nested, out)
	}
	fmt.Println(nested)

	//Example 2:
	appleLaptops := []string{"MacbookPro", "MacbookAir"}
	hpLaptops := []string{"hp650", "hpEliteBook"}
	laptops := [][]string{appleLaptops, hpLaptops}

	for i, v := range laptops {
		fmt.Println("Record: ", i)
		for _, name := range v {
			fmt.Printf("\t Laptop name: %v \n", name)
		}
	}
}

//-----------------------------------------------------------------main function-----------------------------------------------------------------
func main() {
	fmt.Println("------------------------------------------------------------------> Array sample")
	arraySample()

	fmt.Println("------------------------------------------------------------------> Slice sample")
	sliceOperations()
	nestedSlice()
}
