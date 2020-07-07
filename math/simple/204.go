package main

import "fmt"

func main(){
	fmt.Println(countPrimes(10))
}
func countPrimes(n int) int {



	var c int

	if n<=2{
		return 0
	}
	if n == 3{
		return 1
	}
	if n<=4 {
		return 2
	}
	c = 2
	for i:=5;i<n;i++{
		if i%6 != 1 && i%6 != 5{
			continue
		}
		s:=sqrt(i)
		if(s == 2){
			c++
		}
		for j:=3;j<=s;j+=2{
			if(i%j == 0){
				break
			}else if j==s || j+2>s{
				c++
			}
		}
	}
	return c

}

func sqrt(i int)int{
	for j:=1;j<i;j++{
		if j*j==i{
			return j
		}else if j*j>i{
			return j-1
		}
	}
	return i
}
