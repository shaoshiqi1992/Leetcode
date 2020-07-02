package main

import "fmt"

func main(){
	fmt.Println(addBinary("1010","1011"))
}

func addBinary(a string, b string) string {
	// a is always smaller than b
	if len(a) > len(b){
		return addBinary(b,a)
	}
	la := len(a)
	lb := len(b)

	i := la - 1
	j := lb - 1
	var c,d byte
	var r string
	for i >=0 && j >= 0 {

		c = (a[i]-'0')^(b[j]-'0')^d

		if (a[i]-'0')+(b[j]-'0')+d >=2{
			d = 1
		}else{
			d = 0
		}

		if c == 1{
			r ="1"+r
		}else{
			r ="0"+r
		}

		i--; j--;
	}

	if(j>=0){
		for j>=0 {
			c = (b[j]-'0')^d
			d = (b[j]-'0')&d

			if c == 1{
				r="1"+r
			}else{
				r ="0"+r
			}
			j--
		}
	}
	if (d>0){
		r="1"+r
	}

	return string(r)
}