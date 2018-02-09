package gobot

import (
	"errors"
	"strings"
)

// Parse parses a (string) command, and returns a Command struct
func Parse(s string) (Command, error) {
	command := Command{}
	parts := strings.Split(s, " ")
	command.Command = parts[0]
	if len(command.Command) == 0 {
		return command, errors.New("no command detected")
	}
	if len(parts) >= 2 {
		command.Args = strings.Split(parts[1], ",")
	}
	return command, nil
}
