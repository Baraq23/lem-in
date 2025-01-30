package main

import "strings"

// Getslices() organises information from the txt file.
func Getslices(l []string) ([][]string, [][]string, []string, []string) {
	connections := [][]string{}
	points := [][]string{}
	startRoom := []string{}
	endRoom := []string{}

	for i, line := range l {
		if strings.Contains(line, "-") {
			c := strings.Split(line, "-")
			connections = append(connections, c)
		}
		if strings.Contains(line, " ") {
			p := strings.Split(line, " ")
			points = append(points, p)
		}
		if strings.Contains(line, "##start") && i+1 != len(l) {
			startRoom = strings.Split(l[i+1], " ")
		}
		if strings.Contains(line, "##end") && i+1 != len(l) {
			endRoom = strings.Split(l[i+1], " ")
		}
	}
	return connections, points, startRoom, endRoom
}
