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

// optimumPath selects the set of paths with the maximum number of unique paths.
//
// Parameters:
// - setOfPaths: A 3D slice containing multiple sets of 2D slices, 
//   where each 2D slice represents a group of paths.
//
// Returns:
// - The 2D slice (set of paths) with the highest number of paths.
func optimumPaths(setOfPaths [][][]string) [][]string{
	var longestSet [][]string

	for _, pathSlices := range setOfPaths{
		if len(pathSlices) > len(longestSet){
			longestSet = pathSlices
		}
	}
	return longestSet
}

// GetUniquePaths takes a set of paths and identifies the most optimal group of unique paths.
// It ensures that paths in each group do not share intermediate rooms (excluding start and end).
// 
// Parameters:
// - paths: A 2D slice where each inner slice represents a path of rooms.
//
// Returns:
// - A 2D slice representing the group of paths with maximum uniqueness.
//
// Function Overview:
// 1. Iterates over each path as a base reference (`basePath`).
// 2. Marks all rooms in the `basePath` as visited using a map (`visitedRooms`).
// 3. For every other path, checks for uniqueness by ensuring no overlapping intermediate rooms with the `basePath`.
// 4. Appends unique paths (relative to `basePath`) to a group (`uniquePath`).
// 5. Repeats for all paths, storing all groups in `uniquePaths`.
// 6. Uses `optimumPaths` to select and return the group with the maximum unique paths.

func GetUniquePaths(paths [][]string) [][]string{
	var uniquePaths [][][]string

	for i, basePath := range paths{
		uniquePath := make([][]string, 0)
		uniquePath = append(uniquePath, basePath)
		visitedRooms := make(map[string]bool)

		//mark all rooms in the basePath as visited
		for _, room := range basePath{
			visitedRooms[room] = true
		}

		//get unique paths with reference to the base path
		for j, path := range paths{
			if i == j {
				continue // skip the base path
			}

			middleRooms := path[1 : len(path)-1] // exclude start and end rooms
			isUnique := func([]string) bool{
				for _, room := range middleRooms{
					if _, ok := visitedRooms[room]; ok{
						return false
					}
				}
				return true
			}

			if isUnique(path){
				uniquePath = append(uniquePath, path)
				for _, room := range middleRooms{
					visitedRooms[room] = true
				}
			}
		}
		uniquePaths = append(uniquePaths, uniquePath)
	}
return optimumPaths(uniquePaths)
}

// AssignPathsToAnts distributes ants across a set of paths to balance workload.
// Each ant is assigned to the "best" path based on its current capacity.
//
// Parameters:
// - antCount: Total number of ants to be assigned to paths.
// - paths: A slice of paths where each path is a slice of strings (representing rooms).
//
// Returns:
// - A map[int][]string where the key is the ant's ID, and the value is the path assigned to it.
//
// Logic:
// - The function ensures ants are assigned to paths in a way that balances the load. 
//   This is done by considering the path's length and the number of ants already assigned to it.
// - For each ant:
//   - Start by assuming the ant will take the first path (assignedPath = 0).
//   - Iterate through the remaining paths to find the "best" one for the ant.
//     - The "best" path is determined by comparing the length of the paths plus the number of ants already assigned to them.
//     - If a shorter or less-loaded path is found, update `assignedPath`.
//   - Assign the selected path to the current ant in the map `pathsToAnt`.
//   - Increment the count of ants assigned to the chosen path in `antsPerPath`.
//
// Notes:
// - The inner loop uses a greedy approach, exiting early when a better path is not found, ensuring efficiency.
// - The paths are evaluated relative to their lengths and load, promoting fairness in the distribution of ants.
func AssignPathsToAnts(antCount int, paths [][]string) map[int][]string{
	pathsToAnt := make(map[int][]string) // keys are ant IDs and values the paths assigned to each ant
	antsPerPath := make([]int, len(paths)) // number of ants currently assigned to specific path

	for ant := 1; ant <= antCount; ant++{
		assignedPath := 0
		for i := 1; i < len(paths); i++{
			if (len(paths[i-1]) + antsPerPath[i-1]) >= (len(paths[i]) + antsPerPath[i]){
				assignedPath = i
			}else{
				break
			}
		}
		pathsToAnt[ant] = paths[assignedPath] //map current ant to selected path
		antsPerPath[assignedPath]++ // increment number or count of ants assigned to the path 
	}
	return pathsToAnt
}