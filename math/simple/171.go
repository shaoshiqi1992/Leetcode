package main

import "fmt"

func main(){
	fmt.Println(titleToNumber("ZY"))
}
func titleToNumber(s string) int {
	var r,t,n,m int
	for i:=len(s)-1; i>=0 ; i--{
		t = int(s[i] - 'A')+1
		n=m
		for n > 0 {
			t = t*26
			n--
		}
		m++
		r = r+t
	}
	return r
}
