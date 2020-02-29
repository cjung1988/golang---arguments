package main

import (
	"fmt"
	"os"
)

func main() {

	// start the program
	fmt.Println("Start programm")

	// get args from program
	argsWithProg := os.Args

	// print args with prog
	fmt.Printf("%s\n", argsWithProg)

	// get args without prog
	argsWithoutProg := os.Args[1:]

	// print args without prog
	fmt.Printf("%s\n", argsWithoutProg)

	// get length of argument array
	length := len(argsWithoutProg)

	// print length
	fmt.Printf("The arguments array length is %d \n", length)

	// print all data inside of a array
	// for each loop
	for i, s := range argsWithoutProg {
		fmt.Println(i, s)
	}

}
