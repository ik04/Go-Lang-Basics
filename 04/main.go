package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Pizzatron 3000")
	fmt.Println("Rate our pizza")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Your Rating:",input );
	//  := assigns and return the type

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64) // powerful package, gets rid of \n

	if err != nil{
		fmt.Println(err);
		panic(err)
	}else {
		fmt.Println("Added 1 to your rating",numRating + 1);
	}
	
}