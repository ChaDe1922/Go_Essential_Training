//Lets calculate the maximal value of a slice and print it out
package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	//create a holder for the max value and initialize with first value of nums
	max := nums[0]

	//if the current index value is greater than the current max value
	for _, value := range nums[1:] {
		if value > max {
			//replace the current max value with the newest
			max = value
		}
	}
	fmt.Println(max)

}
