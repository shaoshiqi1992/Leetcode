package main

import "fmt"
func main(){
	fmt.Println(isHappy(20))
}
func isHappy(n int) bool {
	var s, i, t int
	var r = map[int]int{}

	s = n
	for s != 1 {
		t = s
		s = 0
		for t > 0 {
			i = t%10
			s += i*i
			t = t/10
		}
		if s == 1{
			return true
		}else if r[s]==1 {
			return false
		}else{
			r[s] = 1
		}

	}
	return true
}