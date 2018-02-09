package gobot

import (
	"errors"
	"fmt"
)

// Grid represents an x,y Grid, such as a Table with 5x5 squares. It's X,Y is always NW=0,0
type Grid struct {
	Height  int
	Width   int
	Objects [][]Object
}

// Init prepares the grids Objects, populating with BlankSquares
func (g *Grid) Init() {
	g.Objects = make([][]Object, g.Height)
	for i := 0; i < g.Height; i++ {
		g.Objects[i] = make([]Object, g.Width)
		for j := 0; j < g.Width; j++ {
			g.Objects[i][j] = BlankSquare{}
		}
	}
}

// MoveObject moves a GridObject to x and y
func (g *Grid) MoveObject(x int, y int, o Object) error {
	lengthOfY := len(g.Objects) - 1
	if lengthOfY < y || y < 0 {
		return fmt.Errorf("that's a bad idea")
	}
	widthOfX := len(g.Objects[y]) - 1
	if widthOfX < x || x < 0 {
		return fmt.Errorf("that's a bad idea")
	}
	cx, cy, _ := g.GetObjectCoordinates(o)
	g.ResetSquare(cx, cy)

	g.Objects[y][x] = o
	return nil
}

// ResetSquare resets a square in the grid to a BlankSquare
func (g *Grid) ResetSquare(x int, y int) {
	g.Objects[y][x] = BlankSquare{}
}

// GetObjectCoordinates returns the X,Y(in its native coordinates) of a GridObject
func (g *Grid) GetObjectCoordinates(o Object) (int, int, error) {
	for y, row := range g.Objects {
		for x, gridObject := range row {
			if gridObject == o {
				return x, y, nil
			}
		}
	}
	return 0, 0, errors.New("404: Object not found")
}
