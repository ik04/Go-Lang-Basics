package main

import "fmt"

func main(){
	fmt.Println("hi mom");
	num := 22
	var ptr = &num
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual value is", *ptr)
	
}