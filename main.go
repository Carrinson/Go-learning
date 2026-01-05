package main

import (
	"log"
//	"sort"
)


func main() {



	// shorthabd way of declaring slices 
	numbers := []int {1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)

	//printing a range of values (in slices)

	log.Println(numbers[4:8])


	// // Slice of strings and integers
	// var mySlice []string
	// var mySliceInt []int
	// mySlice = append(mySlice, "Trevor")
	// mySlice = append(mySlice, "John")

	// mySliceInt = append(mySliceInt, 16)
	// mySliceInt = append(mySliceInt, 18)
	// mySliceInt = append(mySliceInt, 19)
	// mySliceInt = append(mySliceInt, 17)

	// sort.Strings(mySlice)

	// log.Println(mySlice)

	// sort.Ints(mySliceInt)

	// log.Println(mySliceInt)

}
