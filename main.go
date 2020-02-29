package main

import (
	"fmt"
	"os"
)

func main() {

	// start the programm
	fmt.Println("Run programm")

	// get args from programm
	argsWithProg := os.Args

	// print args with prog
	fmt.Printf("%s\n", argsWithProg)

	// get args without prog
	argsWithoutProg := os.Args[1:]

	// print args without prog
	fmt.Printf("%s\n", argsWithoutProg)

	// get lengt of argument array
	length := len(argsWithoutProg)

	// print length
	fmt.Printf("The arguments array length is %d \n", length)

	// print all data inside of a array
	// for each loop
	for i, s := range argsWithoutProg {
		fmt.Println(i, s)
	}
}
