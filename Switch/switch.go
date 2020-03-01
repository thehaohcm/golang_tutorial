package main

import "fmt"

func main(){
	fmt.Println("1. Show \"Anh Hao rat dep trai\"")
	fmt.Println("Please enter a number: ")
	var number int
	fmt.Scan(&number)
	switch number{
		case 1:
			fmt.Println("Anh Hao rat dep trai")
			break;
		case 2:
			fmt.Println("Mr.Hao is very handsome")
		case 3:
			fmt.Println("Both of above 2 sentences are correct")
		default:
			fmt.Println("Please run application and choose again")
	}

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
