package main
import "fmt"
func main(){
	fmt.Println("sum of 2 and 4:",sumWith2Params(2,4))
	fmt.Println("sum of 1, 2 and 3:",sumWith3Params(1,2,3))

	var inputValue int
	fmt.Println("Please enter the number of params you want to calcuate:")
	fmt.Scan(&inputValue)

	var inputNums=make([]int, inputValue)
	for i:=0;i<inputValue;i++{
		fmt.Print("input ",i+1," item value:")
		fmt.Scan(&inputNums[i])
	}
	fmt.Println()

	if len(inputNums)==2{
		fmt.Println("the result is:",sumWith2Params(inputNums[0],inputNums[1]))
	} else if len(inputNums)==3{
		fmt.Println("the result is:",sumWith3Params(inputNums[0],inputNums[1],inputNums[2]))	
	}
}

func sumWith2Params(num1,num2 int) int{
	return num1+num2
}

func sumWith3Params(num1 int, num2 int, num3 int)int {
	return num1+num2+num3
}
