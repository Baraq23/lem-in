package main

//Func Optimize() recieve all paths as a 2d slice the return the optimal paths.
func Optimizer(paths [][]string) [][]string {
	optimized := [][]string{}
	newPaths := paths
	shortPath := []string{}
	for i := 0; i < len(paths); i++ {
		shortPath, newPaths = findShort(newPaths)
		if len(optimized) == 0 {
			optimized = append(optimized, shortPath)
			continue
		}
		if pathsColide(optimized, shortPath) {
			continue
		} else {
			optimized = append(optimized, shortPath)
		}
	}
	return optimized
}

func pathsColide(optimized [][]string, shortPath []string) bool {
	trimed := [][]string{}
	shortPath = shortPath[1 : len(shortPath)-1]
	for _, p := range optimized {
		trimed = append(trimed, p[1:len(p)-1])
	}
	for _, sroom := range shortPath {
		for _, path := range optimized {
			for _, room := range path {
				if sroom == room {
					return true
				}
			}
		}
	}

	return false
}

func findShort(paths [][]string) ([]string, [][]string) {
	short := paths[0]
	ind := 0
	for i, path := range paths {
		if len(short) > len(path) {
			short = path
			ind = i
		}
	}
	paths = append(paths[:ind], paths[ind+1:]...)
	return short, paths
}
