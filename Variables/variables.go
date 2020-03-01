package main

import "fmt"

func main(){
	var a="variable"
	fmt.Println(a)

	var b, c int = 1,2 
	fmt.Print(b," - ",c)

	var d=false
	fmt.Print(" -> ",!d)

	var e int
	fmt.Println("e: ", e)

	f:="apple"
	fmt.Println(f)

	f="d"
	fmt.Println("f value is ",f)
	g:="e"
	fmt.Println("g value is ",g)
}
