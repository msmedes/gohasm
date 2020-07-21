package types

// Command is an Enum for CommandTypes
type Command int

// Command values
const (
	ACommand = iota
	CCommand
	LCommand
	NoCommand
)
