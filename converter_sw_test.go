package gobot

import (
	"testing"
)

func TestSwCoordinateConverterToGrid(t *testing.T) {
	swc := SouthWestConverter{5, 5}
	x, y := swc.ToGridXY(4, 1)
	if x != 4 {
		t.Errorf("x was returned as %d, expected 4", x)
		t.Fail()
	}
	if y != 3 {
		t.Errorf("y was returned as %d, expected 3", y)
		t.Fail()
	}
}

func TestSwCoordinateConverterFromGrid(t *testing.T) {
	swc := SouthWestConverter{5, 5}
	x, y := swc.FromGridXY(4, 3)
	if x != 4 {
		t.Errorf("x was returned as %d, expected 4", x)
		t.Fail()
	}
	if y != 1 {
		t.Errorf("y was returned as %d, expected 1", y)
		t.Fail()
	}
}
