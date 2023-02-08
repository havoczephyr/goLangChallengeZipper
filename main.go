package main

import "fmt"

func main() {
	arrayOne := make([]int, 0)
	arrayTwo := make([]int, 0)

	arrayOne = append(arrayOne, 1, 2, 3)
	arrayTwo = append(arrayTwo, 10, 20, 30)
	zipped := zipArrays(arrayOne, arrayTwo)

	fmt.Println(zipped)
}
