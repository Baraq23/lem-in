package colony

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetNeighbours(inputFile string, antPopulation *AntPopulation, stations *Stations) (Neighbours, map[string][]int, error){
	file, err := os.Open(inputFile)
	if err != nil{
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	neighbours := make(Neighbours)
	rooms := make(map[string][]int)

	scanner := bufio.NewScanner(file)

	var isStart, isEnd bool

	// handle empty file
	if !scanner.Scan(){
		return nil, nil, errors.New("invalid data format, invalid number of Ants")
	}

	// get the total number of ants, which is given as a number in the first line of the file
	text := strings.TrimSpace(scanner.Text())
	if text == ""{
		return nil, nil, errors.New("invalid data format, invalid number of Ants")
	}

	if size, err := strconv.Atoi(text); err == nil{
		antPopulation.Size = size
	}else{
		return nil, nil, errors.New("invalid data format, invalid number of Ants")
	}

	// get the rooms and connections
	for scanner.Scan(){
		text = strings.TrimSpace(scanner.Text())
		if text == ""{
			continue
		}
		if text == "##start"{
			isStart = true
			continue
		}
		if text == "##end"{
			isEnd = true
			continue
		}
		if strings.HasPrefix(text, "#"){
			continue
		}

		// capture the start and end stations
		if isStart{
			stations.Start = strings.Fields(text)[0]
			isStart = false
		}else if isEnd{
			stations.End = strings.Fields(text)[0]
			isEnd = false
		}

		// capture other rooms data
		parts := strings.Fields(text)
		if len(parts) == 3{
			coordinates := make([]int, 2)
			coordinates[0], err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, errors.New("invalid data format, invalid x coordinate")
			}
			coordinates[1], err = strconv.Atoi(parts[2])
			if err != nil{
				return nil, nil, errors.New("invalid data format, invalid y coordinate")
			}
			rooms[parts[0]] = coordinates
			continue
		}

		if isTunnel(parts){
			connection := strings.Split(parts[0], "-")
			if len(connection) != 2{
				return nil, nil, errors.New("invalid data format, invalid tunnel")
			}
			if connection[0] == connection[1]{
				return nil, nil, errors.New("invalid data format, tunnel connects room to itself")
			}

			neighbours[connection[0]] = append(neighbours[connection[0]], connection[1])
			neighbours[connection[1]] = append(neighbours[connection[1]], connection[0])
			continue
		}
		return nil, nil, errors.New("invalid data format")
	}

	if err := scanner.Err(); err != nil{
		return nil, nil, err
	}
	return neighbours, rooms, nil
}

func isTunnel(parts []string) bool{
	return len(parts) == 1 && strings.Contains(parts[0], "-")
}