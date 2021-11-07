package main

import "fmt"

func main() {
	array := []int{4, 76, -29, 3, 223, 9, 0, -3, -44, 76, 98, 3, 62, 144}

	largestRes, _, _ := largest(array...)
	fmt.Printf("largest\n\tmax: %d\n", largestRes)

	alternateRes := alternate(array...)
	fmt.Printf("alternate\n\tmax: %d\n", alternateRes)

	largestTwoRes, largestTwoRes2 := largestTwo(array...)
	fmt.Printf("largestTwo\n\tmax: %d, 2nd: %d\n", largestTwoRes, largestTwoRes2)

	array2 := []int{4, 76, -29, 3, 9, 223, 0, -3, -44, 76, 3, 98, 62, 144}
	sortingTwoRes, sortingTwoRes2 := sortingTwo(array2...)
	fmt.Printf("sortingTwo\n\tmax: %d, 2nd: %d\n", sortingTwoRes, sortingTwoRes2)

	doubleTwoRes, doubleTwoRes2 := doubleTwo(array...)
	fmt.Printf("doubleTwo\n\tmax: %d, 2nd: %d\n", doubleTwoRes, doubleTwoRes2)

	mutableTwoRes, mutableTwoRes2 := mutableTwo(array...)
	fmt.Printf("mutableTwo\n\tmax: %d, 2nd: %d\n", mutableTwoRes, mutableTwoRes2)

	tournamentRes, tournamentRes2 := tournamentTwo(array...)
	fmt.Printf("tournamentTwo\n\tmax: %d, 2nd: %d\n", tournamentRes, tournamentRes2)

}
