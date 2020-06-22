package main

import (
	"fmt"
)

func main(){
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(-121))

}

func isPalindrome(x int) bool {
	if x < 0{
		return false
	}
	var x1 = x
	var r int
	for x1 != 0 {
		var x2 = x1%10
		x1  = x1/10
		r  = r*10+x2
	}
	if r == x {
		return true
	}
	return false
}