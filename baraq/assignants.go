package main

import (
	"fmt"
	"strconv"
)

//Func Assignants() takes in the number of ants and the optimised path and assign each ant of the best path.

func Assignants(n int, optPaths [][]string) {
	//rearrange the  paths with shortest at index 0.
	for i := 0; i < len(optPaths); i++ {
		for j := i+1; j < len(optPaths); j++ {
			if i+1 != len(optPaths) {
				if len(optPaths[i]) > len(optPaths[j]) {
					optPaths[i], optPaths[j] = optPaths[j], optPaths[i]
				}
			}
		}		
	}

	trimmedPaths := [][]string{}
	for i:=0; i < len(optPaths); i++ {
		trimmedPaths = append(trimmedPaths, optPaths[i][1:])
	}

	routeLen := []int{}	
	for i := 0; i < len(trimmedPaths); i++ {
		routeLen = append(routeLen, len(trimmedPaths[i]))
	}
	antRoute := make([][]string, len(routeLen))


	// for i := 1; i <= n; i++ {
		
		//assign first ant
		// if i == 1 {
		// 	antRoute = append(antRoute, []string{"L"+antNum})
		// 	trimmedPaths[0] = append([]string{"X"}, trimmedPaths[0]...)
		// 	routeLen[0]++
		// 	continue
		// }
		antRoute = placeAnt(antRoute, trimmedPaths, routeLen, n)
	// }
	fmt.Println("antroute:\n", antRoute)
	fmt.Println("optPaths:\n", optPaths)
	fmt.Println("trimmedPaths:\n", trimmedPaths)
	fmt.Println("routeLen:\n", routeLen)
}


func placeAnt(antRoute [][]string, trimmedPaths [][]string, routeLen []int, n int ) [][]string {
	if n == 0 {
		return antRoute
	}

	antNum := strconv.Itoa(n)
	
	//check if routLen is leveled
	for i, l := range routeLen {
		if i+1 != len(routeLen) {
			if l < routeLen[i+1] {
				antRoute[i] = append(antRoute[i], "L"+antNum)
				trimmedPaths[i] = append([]string{"X"}, trimmedPaths[i]...)
				routeLen[i]++				
				break
			} else if l > routeLen[i+1] {
				antRoute[i] = append(antRoute[i], "L"+antNum)
				trimmedPaths[i] = append([]string{"X"}, trimmedPaths[i]...)
				routeLen[i]++				
				break
			} 
		}
	}
	return placeAnt(antRoute, trimmedPaths, routeLen, n-1)
}
