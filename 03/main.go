package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to user Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating:")
	// treat problems and errors as variable, handle using comma ok
	//  := the walrus op
	input, _ := reader.ReadString('\n')
	fmt.Println("rating:",input)
}