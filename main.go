package main

import (
	"fmt"
	"speskilltest/helper"
)

func main() {
	fmt.Println("Narcissistic Test, Number (1634) *true")
	result := helper.Narcissistic(1634)
	fmt.Println(result)

	fmt.Println("Narcissistic Test, Number (111) *false")
	result = helper.Narcissistic(111)
	fmt.Println(result)

	fmt.Println("---------------------------------------")

	fmt.Println("Parity Outlier Test, Odd Array [3, 33, 10]")
	arrayInt := []int{3, 33, 10}
	outlierInt, IsOutlier := helper.ParityOutlier(arrayInt)
	if outlierInt != 0 {
		fmt.Println(outlierInt)
	} else {
		fmt.Println(IsOutlier)
	}

	fmt.Println("Parity Outlier Test, Odd Array No Outlier [3, 33, 101]")
	arrayInt = []int{3, 33, 101}
	outlierInt, IsOutlier = helper.ParityOutlier(arrayInt)
	if outlierInt != 0 {
		fmt.Println(outlierInt)
	} else {
		fmt.Println(IsOutlier)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Needle In Haystack Test, Haystack : ['red', 'blue', 'yellow', 'black', 'grey'] , Needle : 'Black'")
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "black"
	index := helper.NeedleInHaystack(haystack, needle)
	fmt.Println(index)

	fmt.Println("---------------------------------------")
	fmt.Println("The Reverse Blue Ocean Test, 1st Array : [10, 10, 35, 42, 54, 80, 10] , 2nd Array : [10]")
	arrayInt1 := []int{10, 10, 35, 42, 54, 80, 10}
	arrayInt2 := []int{10}
	arrayResult := helper.ReverseBlueOcean(arrayInt1, arrayInt2)
	fmt.Println(arrayResult)
}
