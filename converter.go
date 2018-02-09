package gobot

// Converter will convert X,Y coordinates into Grid X,Y coordinates (which speaks NW=0,0)
type Converter interface {
	ToGridXY(x int, y int) (int, int)
	FromGridXY(x int, y int) (int, int)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
