//An even ended number is a number with the same first and last digits
//Question: How many even-ended numbers result from multiplying two four-digit numbers?
//It's easier to check if a number is even ended if you convert the number to the string.package evenEnded
// This can be done by using the func Sprintf (This formats according to a format specifier and returns the resulting string)

package main

import (
	"fmt"
)

//Initialize a counter at 0

func main() {
	counter := 0

	//multiply all 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			num := a * b

			//check if a * b is even ended by first converting the answer to a string
			str := fmt.Sprintf("%d", num)

			//if the first character in the string str is the same as the last charater, this answer is even ended
			if str[0] == str[len(str)-1] {
				//increase the counter by 1 to count the even ended answer
				counter++
			}
		}
	}
	countstr := fmt.Sprintf("%d", counter)
	fmt.Println("There are " + countstr + " even-ended 4-digit multiples")
}
