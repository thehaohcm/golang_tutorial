package main
import "fmt"

func main(){
	fmt.Println("The 1st of loop")
	i:=0
	for i<3 {
		fmt.Println("i=",i)
		i+=1
	}

	fmt.Println("The 2nd of loop")
	for j:=0;j<3;j++{
		fmt.Println("j=",j)
	}

	fmt.Println("Endless loop using for loop")
	brokenNum:=0
	for {
		if (brokenNum==10){
			break;
		}
		fmt.Println("count Number: ",brokenNum)
		brokenNum+=1
	}
}
