package main

import (
	"fmt"
)

func main(){
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(120))

}

func reverse(x int) int {
	if x > 2147483647 || x < -2147483648 {
		return 0
	}

	var x1, r int
	for x != 0{
		x1 = x%10
		x = x/10
		r = r*10+x1
	}

	if r > 2147483647 || r < -2147483648 {
		return 0
	}

	return r
}