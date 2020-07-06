package main
import "fmt"

func main (){
	fmt.Println(trailingZeroes(1808548329));
}
func trailingZeroes(n int) int {
	if n < 5{
		return 0
	}

	return n/5+trailingZeroes(n/5)
}
