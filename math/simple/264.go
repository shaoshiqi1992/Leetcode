package main

import "fmt"

func main(){
	fmt.Println(nthUglyNumber(1352))
}

func nthUglyNumber(n int) int {
	if n == 1{
		return 1
	}
	var i,j int
	for i < n{
		j++
		if ifUgly(j){
			i++
		}
	}
	return j
}

func ifUgly(num int) bool {
	if num == 0{
		return false
	}
	if num == 1{
		return true
	}
	for num%2 == 0 {
		num /=2
	}
	for num%3 == 0{
		num /= 3
	}
	for num%5 == 0{
		num /= 5
	}
	if num == 1{
		return true
	}

	return false
}