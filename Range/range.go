package main
import "fmt"
func main(){
	num:=[5]int{1,2,3,4,5}
	for _,value:= range num{
		fmt.Println("value:",value)
	}

	kvs:=map[int]string{1:"apple", 2:"orange",3:"banana"}

	for k,v:=range kvs{
		fmt.Printf("%d->%s\n",k,v)
	}
}
