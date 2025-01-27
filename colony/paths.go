package colony



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

func contains(slice []string, item string)bool{
	for _, str := range slice{
		if str == item{
			return true
		}
	}
	return false
}