package colony

import (
	"errors"
	"fmt"
)

// CreateAntFarm initializes an AntFarm by creating rooms and assigning ants to paths.
// It takes a map of path-to-ant assignments, along with the start and end room names.
// It creates the necessary rooms, assigns them to the farm, and creates ants with their respective paths.
// Each ant is initialized with its ID, path, and starting position, and the path is converted into a slice of Room objects.
//
// Parameters:
//  - pathToAnt: A map where the key is the ant ID, and the value is a slice of room names representing the ant's path.
//  - start: The name of the starting room.
//  - end: The name of the ending room.
//
// Returns:
//  - *AntFarm: A pointer to the created AntFarm object containing ants, rooms, and the start/end rooms.

func CreateAntFarm(pathToAnt map[int][]string, start, end string) *AntFarm{
	farm := &AntFarm{
		Ants: make([]*Ant, len(pathToAnt)), // slice of pointers to Ant objects, with a length equal to the number of ants
		Rooms: make(map[string]*Room),
	}

	//create rooms
	for _, path := range pathToAnt{
		for _, roomName := range path{
			if _, exists := farm.Rooms[roomName]; !exists{
				room := &Room{Name: roomName}
				if roomName == start{
					room.IsStart = true
					farm.Start = room
				}else if roomName == end{
					room.IsEnd = true
					farm.End = room
				}
				farm.Rooms[roomName] = room
			}
		}
	}
		//create ants and give them paths
		for antID, path := range pathToAnt{
			ant := &Ant{
				Id: antID,
				Path: make([]*Room, len(path)),
				PathIndex: 0,
				CurrentRoom: farm.Start,
				ReachedEnd: false,
			}
			for j, roomName := range path{
				ant.Path[j] = farm.Rooms[roomName]
			}
			farm.Ants[antID-1] = ant
		}
return farm	
}

// validateAnt checks if the provided ant is in a valid state.
// It validates the following:
// 1. The ant is not nil.
// 2. The ant has a valid path (not empty).
// 3. The ant has a current room assigned.
// 4. If the ant is at the start (path length of 2), only one ant can move at a time.
//
// If any of these conditions are violated, it returns an error with a detailed message.
// If the ant is valid, it returns nil.

func(antFarm *AntFarm) validateAnt(ant *Ant) error{
	if ant == nil{
		return errors.New("invalid data format, ant is nil")
	}
	if len(ant.Path) == 0{
		return fmt.Errorf("invalid data format, ant %d has no valid path", ant.Id)
	}
	if ant.CurrentRoom == nil{
		return fmt.Errorf("invalid data format, ant %d has no current room set", ant.Id)
	}
	if len(ant.Path) == 2 && antFarm.Move != 0{
		return fmt.Errorf("only move one ant per turn")
	}
	return nil
}