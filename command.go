package gobot

import (
	"fmt"
	"errors"
	"strconv"
)


// Command represents a command for the robot to process
type Command struct {
	Command string
	Args    []string
}

// Response represents the output from the command. Containing an Error, and words to Speak
type Response struct {
	Error error
	Speak string
}


// HandleCommand handles maps a Command input to a Robot function, executes it, and responds with a CommandResponse
func (r *Robot) HandleCommand(cmd Command) Response {
	cr := Response{}
	switch cmd.Command {
	case "PLACE":
		if len(cmd.Args) != 3 {
			cr.Error = errors.New("Malformed place command")
			return cr
		}
		x, _ := strconv.Atoi(cmd.Args[0])
		y, _ := strconv.Atoi(cmd.Args[1])
		err := r.Place(x, y, cmd.Args[2])
		if err != nil {
			cr.Error = err
			return cr
		}
		r.Placed = true
		return cr
	}

	if r.Placed == false {
		cr.Error = fmt.Errorf("Must be PLACEd first")
		return cr
	}

	switch cmd.Command {
	case "MOVE":
		cr.Error = r.Move()
	case "LEFT":
		cr.Error = r.TurnLeft()
	case "RIGHT":
		cr.Error = r.TurnRight()
	case "REPORT":
		report, err := r.Report()
		cr.Speak = report
		cr.Error = err
	default:
		cr.Speak = fmt.Sprintf("%s the gobot can not compute command: %s", r.Name, cmd.Command)
	}
	return cr
}
