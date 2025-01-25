package colony

type Ant struct{
	Id int
	Path []*Room
	PathIndex int
	CurrentRoom *Room
	ReachedEnd bool
}

type Room struct{
	Name string
	IsStart bool
	IsEnd bool
}

type Farm struct{
	Ants []*Ant
	Move int
	Rooms map[string]*Room
	Start *Room
	End *Room
}

type Path struct{
	Rooms []*Room
}