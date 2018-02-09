package gobot

import (
	"testing"
)

func TestParseSimpleCommand(t *testing.T) {
	i := "SIMPLE"
	command, err := Parse(i)
	if err != nil {
		t.Errorf("Error returned while parsing simple command %s", err)
	}
	if command.Command != i {
		t.Errorf("Expected parser to return Command %s, but instead returned %s", i, command.Command)
	}
}

func TestParseCommandReturnsArgs(t *testing.T) {
	i := "SIMPLE X,Y,Z"
	command, err := Parse(i)
	if err != nil {
		t.Errorf("Error returned while parsing simple command %s", err)
		t.Fail()
		return
	}
	if len(command.Args) != 3 {
		t.Errorf("Expected Args to have a len() of 2, got %b", len(command.Args))
		t.Fail()
		return
	}
	if command.Args[0] != "X" {
		t.Errorf("Arg 0 was not X, it was %s", command.Args[0])
		t.Fail()
	}
	if command.Args[1] != "Y" {
		t.Errorf("Arg 1 was not Y, it was %s", command.Args[1])
		t.Fail()
	}
	if command.Args[2] != "Z" {
		t.Errorf("Arg 2 was not Z, it was %s", command.Args[2])
		t.Fail()
	}
}

func TestEmptyCommandInvalid(t *testing.T) {
	_, err := Parse("")
	if err == nil {
		t.Errorf("Expected empty command to be invalid")
	}
}
