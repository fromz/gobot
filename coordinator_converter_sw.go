package gobot

// SouthWestConverter converts X,Y coordinates from SW=0,0 to NW=0,0
type SouthWestConverter struct {
	Height int //y
	Width  int //x
}

// ToGridXY converts X,Y from SW=0,0 [robot] coordinates to NW=0,0 [grid]
func (swc SouthWestConverter) ToGridXY(x int, y int) (int, int) {
	return x, abs(y - (swc.Height - 1))
}

// FromGridXY converts X,Y from NW=0,0 [grid] to SW=0,0 [robot]
func (swc SouthWestConverter) FromGridXY(x int, y int) (int, int) {
	return x, abs(y - (swc.Height - 1))
}
