package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)
var params []string
var datas = map[int]string{}

func main(){
	var res [][]int
	readfile()
	//fmt.Println(params,datas)
	params1,_ := strconv.Atoi(params[0])
	params2,_ := strconv.Atoi(params[1])

	for j:=0;j < params1;j++{
		res = append(res,[]int{})
	}

	for k:=0;k <params2;k ++{
		assign := 0
		for j:=0;j < params1;j++{
			if(cal(res[j])<cal(res[assign])){
				assign = j
			}
		}
		t,_ := strconv.Atoi(datas[k])
		res[assign] = append(res[assign],t)
	}

	max := cal(res[0])
	for j:=0;j < params1;j++{
		if(cal(res[j])>max){
			max = cal(res[j])
		}
	}

	fmt.Println(max)
}
func cal(a []int) int{
	var b int
	for _,value := range a{
		b += value
	}
	return b
}
func readfile() {
	// 读取一个文件的内容
	file, err := os.Open("approxiate/3.txt")
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 处理结束后关闭文件
	defer file.Close()

	// 使用bufio读取
	r := bufio.NewReader(file)
	n := 0
	for {
		// 分行读取文件  ReadLine返回单个行，不包括行尾字节(\n  或 \r\n)
		data, _, err := r.ReadLine()

		// 读取到末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}

		// 打印出内容
		if n == 0{
			params=strings.Fields(string(data))
		}else{
			datas[n-1] = string(data)
		}
		n++
	}
}