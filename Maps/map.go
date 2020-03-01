package main
import "fmt"
func main(){
	m:=make(map[string]int)

	m["k1"]=4
	m["k2"]=45

	fmt.Println("map:",m)

	v1:=m["k1"]
	fmt.Println("v1=",v1)

	fmt.Println("len:",len(m))

	//delete the specified element
	delete(m,"k1")
	fmt.Println("map after delete element:",m)

	_,prs:=m["k1"]
	fmt.Println("value of k1 after deleted:",prs)

	_,prs=m["k2"]
	fmt.Println("value of k2:",prs)

	k2Value:=m["k2"]
	fmt.Println("value of k2 (actual):",k2Value)
}
