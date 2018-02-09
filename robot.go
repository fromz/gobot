package gobot

import (
	"fmt"
	"errors"
)

// Robot stores details of the Grid (table) he is placed on, and other information
type Robot struct {
	Name                string
	Icon                string
	Placed              bool
	Direction           string
	Grid                Grid
	CoordinateConverter Converter
}

// DecorateSquare returns a string containing an ASCII representation of this robot
func (r Robot) DecorateSquare() string {
	var decoration string
	switch r.Direction {
	case "NORTH":
		decoration = "⇑"
		break
	case "SOUTH":
		decoration = "⇓"
		break
	case "EAST":
		decoration = "⇒"
		break
	case "WEST":
		decoration = "⇐"
		break
	}
	return decoration + r.Icon
}

// IsValidDirection takes a string and returns true if this is a valid direction, and false if invalid
func (r *Robot) IsValidDirection(direction string) bool {
	return direction == "NORTH" || direction == "SOUTH" || direction == "WEST" || direction == "EAST"
}

// HasValidDirection returns true if the robot currently has a valid direction, and false if invalid
func (r *Robot) HasValidDirection() bool {
	return r.IsValidDirection(r.Direction)
}

// GetCoordinates returns the robots current X,Y coordinate from its (SW=0,0) perspective
func (r *Robot) GetCoordinates() (int, int, error) {
	x, y, err := r.Grid.GetObjectCoordinates(r)
	nx, ny := r.CoordinateConverter.FromGridXY(x, y)
	return nx, ny, err
}

// HasValidCoordinates returns true if the robot has valid coordinates, and false if invalid
func (r *Robot) HasValidCoordinates() bool {
	_, _, err := r.GetCoordinates()
	if err != nil {
		return false
	}
	return true
}

// IsPlaced returns true if the robot is currently placed on the table, and false if not
func (r *Robot) IsPlaced() bool {
	if r.HasValidCoordinates() == false {
		return false
	}
	return true
}

// Report produces a (string) report of the robots coordinates and direction, and an error if invalid direction or coordinates
func (r *Robot) Report() (string, error) {
	x, y, err := r.GetCoordinates()
	if err != nil {
		return "", err
	}
	if r.HasValidDirection() == false {
		return "", fmt.Errorf("Could not produce report, as %s has an invalid direction of %s", r.Name, r.Direction)
	}
	return fmt.Sprintf("%d,%d,%s", x, y, r.Direction), nil
}

// TurnLeft rotates robot by 90 degrees left
func (r *Robot) TurnLeft() error {
	switch r.Direction {
	case "NORTH":
		r.Direction = "WEST"
		break
	case "SOUTH":
		r.Direction = "EAST"
		break
	case "EAST":
		r.Direction = "NORTH"
		break
	case "WEST":
		r.Direction = "SOUTH"
		break
	default:
		return errors.New("No direction set")
	}
	return nil
}

// TurnRight rotates robot by 90 degrees right
func (r *Robot) TurnRight() error {
	switch r.Direction {
	case "NORTH":
		r.Direction = "EAST"
		break
	case "SOUTH":
		r.Direction = "WEST"
		break
	case "EAST":
		r.Direction = "SOUTH"
		break
	case "WEST":
		r.Direction = "NORTH"
		break
	default:
		return errors.New("No direction set")
	}
	return nil
}

// Move moves the robot by one in its current direction, if it wont fall
func (r *Robot) Move() error {
	x, y, err := r.GetCoordinates()
	nx, ny := r.CoordinateConverter.ToGridXY(x, y)
	if err != nil {
		return err
	}
	switch r.Direction {
	case "NORTH":
		ny = ny - 1
	case "SOUTH":
		ny = ny + 1
	case "EAST":
		nx = nx + 1
	case "WEST":
		nx = nx - 1
	default:
		return errors.New("No direction set. Can not move.")
	}
	err = r.Grid.MoveObject(nx, ny, r)
	if err != nil {
		return err
	}

	return nil
}

// Place positions the robot to an X/Y coordinate facing a particular direction
func (r *Robot) Place(x int, y int, direction string) error {
	if r.IsValidDirection(direction) == false {
		return fmt.Errorf("%s is an invalid direction", direction)
	}
	nx, ny := r.CoordinateConverter.ToGridXY(x, y)
	err := r.Grid.MoveObject(nx, ny, r)
	if err != nil {
		return err
	}
	r.Direction = direction
	return nil
}
