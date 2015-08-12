package main

import "math"

type N009PalindromeNumber struct {
}

func (this *N009PalindromeNumber) isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	length := 1
	temp := x
	temp2 := 0
	temp /= 10
	for temp > 0 {
		temp /= 10
		length++
	}
	for i := length; i > 1; i = i - 2 {
		temp2 = int(math.Pow(10, float64(i-1)))
		temp = x / temp2
		if temp != x%10 {
			return false
		}
		if temp != 0 {
			x = (x - temp*temp2) / 10
		} else {
			x /= 10
		}
	}
	return true
}
