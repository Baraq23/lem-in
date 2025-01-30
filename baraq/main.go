package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [Name-of-file.txt]")
		fmt.Println("Example: go run . <fileName.txt>")
		return
	}
	inputFile := os.Args[1]
	if ext := filepath.Ext(inputFile); ext != ".txt" {
		fmt.Println("only files with .txt extension are accepted")
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error Opening file: ", err)
		return
	}
	defer file.Close()

	// get file content into slice
	fileSlice := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		fileSlice = append(fileSlice, line)
	}

	numberofants, err := strconv.Atoi(fileSlice[0])
	if err != nil {
		fmt.Println("ERROR: invalid number of Ants")
		return
	}
	fmt.Sprintf("Number of ants: %t\n", numberofants)

	connections, points, startRoom, endRoom := Getslices(fileSlice)
	// fmt.Println(connections, points, startRoom, endRoom)
	allPaths, err := DFS(connections, points, startRoom, endRoom)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print all paths
	fmt.Printf("All possible paths from %s to %s:\n", string(startRoom[0]), string(endRoom[0]))
	for _, path := range allPaths {
		fmt.Println(path)
	}

	optimizedPaths := Optimizer(allPaths)

	fmt.Printf("Optimized paths from %s to %s:\n", string(startRoom[0]), string(endRoom[0]))
	for _, path := range optimizedPaths {
		fmt.Println(path)
	}
}

