package main

import "fmt"
func main(){
	 fmt.Println(convertToTitle(1));
}

func convertToTitle(n int) string {
	alpha := [] string{
		"Z","A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z"}

	var c, d int
	var res string
	for c = n;c > 0 ; c=c/26{
		d = c%26
		if(d == 0){
			c = c - 1
		}
		res = alpha[d] + res
	}
	return res

}