package main

import (
	"strconv"
)

// Func Assignants() takes in the number of ants and the optimised path and assign each ant of the best path.
func Assignants(n int, optPaths [][]string) [][]string{
	// rearrange the  paths with shortest at index 0.
	for i := 0; i < len(optPaths); i++ {
		for j := i + 1; j < len(optPaths); j++ {
			if i+1 != len(optPaths) {
				if len(optPaths[i]) > len(optPaths[j]) {
					optPaths[i], optPaths[j] = optPaths[j], optPaths[i]
				}
			}
		}
	}

	trimmedPaths := [][]string{}
	for i := 0; i < len(optPaths); i++ {
		trimmedPaths = append(trimmedPaths, optPaths[i][1:])
	}

	routeLen := []int{}
	for i := 0; i < len(trimmedPaths); i++ {
		routeLen = append(routeLen, len(trimmedPaths[i]))
	}
	antRoute := make([][]string, len(routeLen))

	antRoute = placeAnt(antRoute, trimmedPaths, routeLen, n)
	return antRoute
}

func placeAnt(antRoute [][]string, trimmedPaths [][]string, routeLen []int, n int) [][]string {
	if n == 0 {
		return antRoute
	}

	antNum := strconv.Itoa(n)

	// check if routLen is leveled
	low := routeLen[0]
	indx := 0
	for i, l := range routeLen {
		if l < low {
			indx = i
		}
	}
	antRoute[indx] = append(antRoute[indx], "L"+antNum)
	trimmedPaths[indx] = append([]string{"X"}, trimmedPaths[indx]...)
	routeLen[indx]++

	return placeAnt(antRoute, trimmedPaths, routeLen, n-1)
}
