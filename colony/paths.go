package colony


// GetAllPaths finds all possible paths from a starting room (`start`) to an ending room (`end`) 
// in a graph represented as a `Neighbours` map. Each key in the map represents a room, 
// and the value is a slice of its neighboring rooms.
//
// Parameters:
// - neighbours: A map where each key is a room name (string) and its value is a slice of neighboring room names.
// - start: The name of the starting room.
// - end: The name of the ending room.
//
// Returns:
// - A 2D slice of strings, where each inner slice represents a unique path from the start to the end room.
//
// How it works:
// 1. The `paths` variable stores all the valid paths found.
// 2. A nested function `depthSearch` performs a Depth-First Search (DFS):
//    - It appends the current room to the path.
//    - If the current room matches the `end` room, it saves a copy of the path to `paths` and returns.
//    - Otherwise, it iterates through all neighbors of the current room and recursively visits them,
//      ensuring that no room is visited more than once in the current path (using `contains`).
// 3. `depthSearch` is initially called with the `start` room and an empty path.
// 4. The function returns all collected paths once the DFS completes.

func GetAllPaths(neighbours Neighbours, start, end string)[][]string{
	var paths [][]string
	var depthSearch func(string, []string)

	depthSearch = func(currentRoom string, path []string){
		path = append(path, currentRoom)

		if currentRoom == end {
			paths = append(paths, append([]string(nil), path...))
			return
		}

		for _, neighbour := range neighbours[currentRoom]{
			if !contains(path, neighbour){
				depthSearch(neighbour, path)
			}
		}
	}
	depthSearch(start, []string{})
	return paths
}

// contains checks if a given string `item` exists in a slice of strings `slice`.
// Returns:
// - true if `item` is found in `slice`, false otherwise.

func contains(slice []string, item string)bool{
	for _, str := range slice{
		if str == item{
			return true
		}
	}
	return false
}