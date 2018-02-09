package gobot

// Object represents anything that's on the grid (including blank grids)
type Object interface {
	DecorateSquare() string
}

// BlankSquare is exactly that. A Blank Square for use in a Grid.
type BlankSquare struct {
}

// DecorateSquare decorates the square with an ASCII space
func (r BlankSquare) DecorateSquare() string {
	return " "
}
