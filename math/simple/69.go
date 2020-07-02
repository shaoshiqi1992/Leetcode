package main

import "fmt"
func main(){
	fmt.Println(mySqrt(6))
}

func mySqrt(x int) int {
	if x == 1{
		return 1
	}
	y := x/2
	for y*y > x{
		y = y/2
	}
	for i:= y; i<= 2*y; i++ {

		if i*i > x{
			y = i-1
			break
		}
		y = i
	}

	return y

}