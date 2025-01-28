package colony

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// DisplayOutput is the main function to process input, compute ant movements, and display results.
// It reads a file containing data for an ant population and room stations, processes paths, and simulates ant movements.
//
// Function Steps:
// 1. Reads the file path from command-line arguments (`os.Args[1]`):
//    - Validates that the file has a `.txt` extension.
//    - If the file extension is invalid, it prints an error message and exits.
// 2. Initializes structures for ant population and room stations (`AntPopulation` and `Stations`).
// 3. Retrieves the neighboring rooms using `GetNeighbours`:
//    - If an error occurs during this process, it prints the error and exits.
// 4. Extracts the start and end stations from `stations`.
// 5. Computes all paths from the start to the end using `GetAllPaths`.
// 6. Sorts the paths in ascending order of length using `sort.SliceStable`:
//    - Ensures paths of equal length retain their original relative order.
// 7. Filters out duplicate paths with `GetUniquePaths`.
// 8. Assigns paths to ants using `AssignPathToAnt` based on the population size and unique paths.
// 9. Creates an `AntFarm` with the assigned paths and the start and end rooms.
// 10. Simulates ant movements using `farm.SimulateMovement`:
//     - If an error occurs during the simulation, it prints the error and exits.
// 11. Reads the input file content and prints it to the console:
//     - Handles potential file read errors gracefully.
// 12. Prints the simulation results (`moves`) to the console.
func DisplayOutput(){
	inputFile := os.Args[1]
	fileExtension := filepath.Ext(inputFile)

	if fileExtension != ".txt"{
		fmt.Printf("only files with .txt extension are accepted")
		return
	}
	antPopulation := &AntPopulation{}
	stations := &Stations{}

	// get the adjacent rooms for each room
	neighbours, _, err := GetNeighbours(inputFile, antPopulation, stations)
	if err != nil{
		fmt.Println("ERROR: ", err)
		return
	}
	start := stations.Start
	end := stations.End
	paths := GetAllPaths(neighbours, start, end)

	// sort in ascending order, in a stable manner(elements with same length remains in their original position)
	sort.SliceStable(paths, func(i, j int) bool{
		return len(paths[i]) < len(paths[j])
	})

	uniquePaths := GetUniquePaths(paths)
	antAssignments := AssignPathToAnt(antPopulation.Size, uniquePaths)

	farm := CreateAntFarm(antAssignments, start, end)

	moves, err := farm.SimulateMovement()
	if err != nil{
		fmt.Println("ERROR: ", err)
		return
	}


	fileContent, err := os.ReadFile(inputFile)
	if err != nil{
		fmt.Printf("error reading file: %s\n", err)
		return
	}
	fmt.Println(string(fileContent))

	fmt.Println()
	fmt.Print(moves)
}

