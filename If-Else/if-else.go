package main

import (
	"fmt"
//	"os"
//	"bufio"
)

func main(){
//	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("Please input a number: ")
	var num int
	fmt.Scan(&num)
	result:=num%2

	if result==0 {
		fmt.Println(num,"is a even number")
	}else{
		fmt.Println(num,"is a odd number")
	}
}
