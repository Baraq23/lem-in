package colony

import (
	"errors"
	"fmt"
	"strings"
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

//allAntsReached checks if all ants in the AntFarm have reached the end room
func(antFarm *AntFarm) allAntsReached() (bool, error){
	for _, ant := range antFarm.Ants{
		if err := antFarm.validateAnt(ant); err != nil{
			return false, err
		}
		if !ant.ReachedEnd{
			return false, nil
		}

	}
	return true, nil
}

// moveAnt handles the movement of a single ant along its assigned path.
//
// Parameters:
// - ant: A pointer to the Ant that is moving.
// - occupiedRooms: A map where keys are Room pointers and values are the Ant occupying each room.
//
// Returns:
// - A formatted string representing the move in the format "L<ant_id>-<room_name>".
// - An empty string if the ant cannot move (e.g., it has reached the end of its path or the next room is occupied).
func(antFarm *AntFarm) moveAnt(ant *Ant, occupiedRooms map[*Room]*Ant)string{
	if ant.PathIndex >= len(ant.Path)-1 {
		return ""
	}

	nextRoom := ant.Path[ant.PathIndex+1]

	//check availability of next room
	if occupiedRooms[nextRoom] != nil && !nextRoom.IsEnd{
		return ""
	}
	//clear room if it's not start or end 
	if !ant.CurrentRoom.IsStart && !ant.CurrentRoom.IsEnd{
		occupiedRooms[ant.CurrentRoom] = nil
	}
	// move ant to next room
	ant.CurrentRoom = nextRoom
	ant.PathIndex++
	//mark the room as occupied if it is not the start or end
	if !nextRoom.IsStart && !nextRoom.IsEnd{
		occupiedRooms[nextRoom] = ant
	}
	//check if the end room has been reached
	if nextRoom.IsEnd{
		ant.ReachedEnd = true
	}

	return fmt.Sprintf("L%d-%s", ant.Id, nextRoom.Name)
}

// getAntMoves calculates and performs the movements of all ants in the AntFarm.
// It clears the occupancy of non-start and non-end rooms, validates each ant, and processes their movement.
// 
// Parameters:
// - occupiedRooms (map[*Room]*Ant): A map tracking which ants occupy which rooms. Keys are Room pointers, and values are Ant pointers.
//
// Returns:
// - []string: A slice of strings describing the movements of ants during the function call.
//
// Function Details:
// 1. Clears the `occupiedRooms` map for all rooms except the start and end rooms to reset the state for the next round.
// 2. Iterates through all ants in the AntFarm and performs the following:
//    - Validates each ant using the `validateAnt` method. If validation fails, skips to the next ant.
//    - Skips ants that have already reached the end room.
//    - Moves the ant to its next position using the `moveAnt` method, which updates the `occupiedRooms` map accordingly.
//    - Increments the AntFarm's `Move` counter if the ant has a direct path consisting of only the start and end rooms.
//    - Appends the movement description to the `moves` slice if the ant successfully moves.
// 3. Returns the `moves` slice, which contains all successful movement descriptions for the current function call.
func (antFarm *AntFarm) getAntMoves(occupiedRooms map[*Room]*Ant) []string{
	var moves []string

	// clear occupied rooms except for start and end
	for room := range occupiedRooms{
		if !room.IsStart && !room.IsEnd{
			occupiedRooms[room] = nil
		}
	}

	for _, ant := range antFarm.Ants{
		if err := antFarm.validateAnt(ant); err != nil{
			continue // skip to the next ant if there is an error
		}

		if ant.ReachedEnd{
			continue //skip if the ant has reached the end
		}

		move := antFarm.moveAnt(ant, occupiedRooms)
		if len(ant.Path) == 2{
			antFarm.Move++
		}
		if move != ""{
			moves = append(moves, move)
		}
	}
	return moves
}

// SimulateMovement orchestrates the entire ant movement simulation within the AntFarm.
// It processes the movements of ants step-by-step until all ants have reached their destinations.
//
// Returns:
// - string: A formatted string containing all moves made during the simulation, with each step's moves on a new line.
// - error: An error if the simulation cannot proceed, such as when there are no ants.
//
// Function Details:
// 1. Checks if the AntFarm contains ants (`len(antFarm.Ants) == 0`):
//    - If no ants are found, it returns an error indicating invalid data.
// 2. Initializes a map `occupiedRooms` to track which ants occupy which rooms during the simulation.
// 3. Declares a slice `simulatedMoves` to record all movements of ants across all steps.
// 4. Retrieves the status of whether all ants have reached their destination (`allAntsReached`) using `antFarm.allAntsReached()`.
// 5. Enters a loop that continues as long as not all ants have reached their destination and no error occurred:
//    - Calls `getAntMoves` to process the current step's movements.
//    - Appends the joined movements for the step to `simulatedMoves` if there are any moves in the current step.
//    - Resets the `Move` counter in `antFarm` to `0` after processing the step (if non-zero).
// 6. Exits the loop when all ants have reached their destination or an error occurs.
// 7. Returns the joined `simulatedMoves` as a single string, with each step's moves separated by a newline, and a `nil` error.

func (antFarm *AntFarm) SimulateMovement()(string, error){
	if len(antFarm.Ants) == 0{
		return "", errors.New("invalid data format, invalid number of ants")
	}

	occupiedRooms := make(map[*Room]*Ant)
	var simulatedMoves []string

	allAntsReached, err := antFarm.allAntsReached()

	for !allAntsReached && err == nil{
		moves := antFarm.getAntMoves(occupiedRooms)
		if len(moves) > 0{
			simulatedMoves = append(simulatedMoves, strings.Join(moves, " "))
		}
		if antFarm.Move != 0{
			antFarm.Move = 0
		}
	}
return strings.Join(simulatedMoves, "\n") + "\n", nil
}