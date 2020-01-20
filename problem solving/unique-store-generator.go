package main

import (
	"fmt"
	"math/rand"
)

// SetDefault : set default value to uniqueIDGrid
func SetDefault(rows int, cols int, grid [][]int, uniqueIDGrid [][]string) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				uniqueIDGrid[i][j] = "0"
			} else {
				uniqueIDGrid[i][j] = "*"
			}
		}
	}
}

// PrintAry : print ary element
func PrintAry(rows int, cols int, uniqueIDGrid [][]string) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s ----------- ", uniqueIDGrid[i][j])
		}
		fmt.Println("")
	}
}

// Update : updates the uniqueIDGrid
func Update(i int, j int, k int, l int, uniqueIDGrid [][]string) bool {
	if uniqueIDGrid[k][l] == "0" || uniqueIDGrid[k][l] == "*" {
		return false
	}
	uniqueIDGrid[i][j] = uniqueIDGrid[k][l]
	return true
}

// Find : finds given item exists in slice
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindTotalDistinct : finds distinct number of stores
func FindTotalDistinct(rows int, cols int, uniqueIDGrid [][]string) int {

	starCount := 0
	var items []string

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if uniqueIDGrid[i][j] == "0" {
				continue
			}

			if uniqueIDGrid[i][j] == "*" {
				starCount++
			} else {
				_, found := Find(items, uniqueIDGrid[i][j])
				if !found {
					items = append(items, uniqueIDGrid[i][j])
				}

			}
		}
	}
	return starCount + len(items)
}

// TotalUniqueStores : calculates the total nummber of unique stores can be generated
func TotalUniqueStores(rows int, cols int, grid [][]int) int {

	// init uniqueIDGrid slice
	var uniqueIDGrid [][]string
	for i := 0; i < rows; i++ {
		r := make([]string, cols)
		for j := 0; j < cols; j++ {
			r[j] = "*"
		}
		uniqueIDGrid = append(uniqueIDGrid, r)
	}

	SetDefault(rows, cols, grid, uniqueIDGrid)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Skip processing those are 0 or already processed
			if uniqueIDGrid[i][j] == "0" || uniqueIDGrid[i][j] != "*" {
				continue
			}

			// Left
			if j-1 >= 0 && grid[i][j-1] == 1 {
				var passed = Update(i, j, i, j-1, uniqueIDGrid)
				if passed {
					continue
				}
			}

			// Right
			if j+1 < cols && grid[i][j+1] == 1 {
				var passed = Update(i, j, i, j+1, uniqueIDGrid)
				if passed {
					continue
				}
			}

			// Up
			if i-1 >= 0 && grid[i-1][j] == 1 {
				var passed = Update(i, j, i-1, j, uniqueIDGrid)
				if passed {
					continue
				}
			}

			// Down
			if i+1 < rows && grid[i+1][j] == 1 {
				var passed = Update(i, j, i+1, j, uniqueIDGrid)
				if passed {
					continue
				}
			}

			// Still not updated means no pair found
			if uniqueIDGrid[i][j] == "*" {
				uniqueIDGrid[i][j] = string(rand.Intn(100000))
			}
		}
	}

	PrintAry(rows, cols, uniqueIDGrid)
	return FindTotalDistinct(rows, cols, uniqueIDGrid)
}

func main() {
	r := 4
	c := 3
	a := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{0, 0, 0},
		{1, 1, 1},
	}

	x := TotalUniqueStores(r, c, a)
	fmt.Println("Total store: ", x)

	r1 := 5
	c1 := 5
	a1 := [][]int{
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	x1 := TotalUniqueStores(r1, c1, a1)
	fmt.Println("Total store: ", x1)

}
