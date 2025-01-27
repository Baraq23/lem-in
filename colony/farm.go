package colony

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