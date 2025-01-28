package main

import (
	"fmt"
	"lem-in/colony"
	"os"
)


func main(){
	if len(os.Args) != 2{
		fmt.Println("Error: Invalid number of arguments.\nUsage: go run . <inputFile.txt>")
		return
	}
	colony.DisplayOutput()
}