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

//-----------------------------------------------------------------map (dictionary type)--------------------------------------------------------
func mapOperations() {
	// defining a map in 3 ways
	//var sampleMap = map[string]int
	//sampleMap := map[string]int

	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	//a. Adding to the map:
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	//b. Remove from the map:
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	//c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	//Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}

	// item contains checking
	if value, ok := currency["USD"]; ok { //comma ok idiom
		fmt.Printf("The value %s is present\n", value)
		fmt.Println(ok)
	}

}

func nestedMap() {
	currency := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	for key, value := range currency {
		fmt.Printf("Currency Name: %v\n", key)
		for k, v := range value {
			fmt.Printf("\t Currency Code: %v\n\t\t\t Ranking: %v\n\n", k, v)
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

	fmt.Println("------------------------------------------------------------------> Map sample")
	mapOperations()
	nestedMap()
}
