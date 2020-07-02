package main

import "fmt"

func main(){
	fmt.Println(romanToInt("IX"))
}

func romanToInt(s string) int {
	var alpha = map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000}
	var r int
	for i:=0 ; i<len(s) ; i++ {
		if i < len(s)-1 && (
			(string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X")) ||
			(string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C")) ||
			(string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M"))){
			r+=(alpha[string(s[i+1])]-alpha[string(s[i])])
			i++
		}else{
			r += alpha[string(s[i])]
		}
	}

	return r
}
