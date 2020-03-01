package main
import (
	"fmt"
	"math/rand"
	"time"
		)

func main(){
	var a[5] int
	fmt.Println("emp=",a)

	b:=[3] int{1,2}

	for i:=0;i<3;i++ {
		fmt.Println("value ",i+1, ":",b[i])
	}

	fmt.Println("2Dimensional array")
	rand.Seed(time.Now().UnixNano())
	var c[5][6] int
	for i:=0;i<5;i++ {
		for j:=0;j<6;j++ {
			c[i][j]=rand.Intn(100)
			fmt.Print("    ",c[i][j])
		}
		fmt.Println()
	}
}
