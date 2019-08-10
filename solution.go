package main

import "fmt"

func isDigitsOdd(x int) bool {
	for {
		if x < 10 {
			if x%2 == 0 {
				return false
			} else {
				return true
			}
		} else {
			var currentDigit = x % 10
			x = x / 10
			if currentDigit%2 == 0 {
				return false
			}
		}
	}
}

func reverse(x int) int {
	// continually remove last digit from original
	// add to end of new_int
	new_int := 0
	for x > 0 {
		remainder := x % 10
		new_int *= 10
		new_int += remainder
		x /= 10
	}
	return new_int
}

func hasTrailingZero(x int) bool {
	if x%10 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var count = 0
	for i := 1; i <= 100000000; i++ {
		if hasTrailingZero(i) == false {
			var sum = i + reverse(i)
			if isDigitsOdd(sum) {
				count++
			}
		}
	}
	fmt.Println(count)
}
