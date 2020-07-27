package main

import "fmt"
func main(){
	fmt.Println(isUgly(1808548329))
}
func isUgly(num int) bool {
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

