//Beginning Challenge: Fizz Buzz
//This solution works, but there are other ways to do it as well!
package main

import "fmt"

func main() {
	//Print numbers 1 - 20
	for num := 0; num < 21; num++ {
		if num < 1 {
			continue
		}

		//If number is divisible by 3 and not divisible by 5, print "fizz"
		if num%3 == 0 && num%5 != 0 {
			fmt.Println("fizz")
			continue
		}

		//if Number is divisible by 5 and not divisible by 3, print "buzz"
		if num%5 == 0 && num%3 != 0 {
			fmt.Println("buzz")
			continue
		}

		//if number is divisible by both 3 & 5, print "fizzbuzz"
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		fmt.Println(num)

	}

}
