package parser

import (
	"fmt"
	"gohasm/internal/pkg/types"
)

type CInstruction struct {
	comp string
	dest string
	jump string
}

type Parser struct {
	filepath           string
	LineNumber         int
	file               []string
	currentCommandType types.Command
	symbol             string
	comp               string
	dest               string
	jump               string
	InstructionCounter int
}

func NewParser(filepath string) *Parser {
	return &Parser{
		filepath:           filepath,
		LineNumber:         -1,
		file:               loadFile(),
		InstructionCounter: 0,
	}
}

func loadFile() []string {
	return []string{"ok"}
}

func (p *Parser) Reset() {
	p.LineNumber = 0
}

func (p *Parser) resetCommands() {
	p.currentCommandType = types.NoCommand
	p.symbol = ""
	p.comp = ""
	p.dest = ""
	p.jump = ""
}

func (p *Parser) HasMoreCommands() bool {
	return p.LineNumber < len(p.file)-1
}

func (p *Parser) Advance() {
	p.LineNumber++
	p.currentCommandType = p.commandType()
	if p.currentCommandType == types.ACommand {
		p.symbol = p.parseACommand()
	}
	if p.currentCommandType == types.CCommand {
		instruction := p.parseCCommand()
		p.dest = instruction.dest
		p.comp = instruction.comp
		p.jump = instruction.jump
	}
	if p.currentCommandType == types.LCommand {
		p.symbol = p.parseLCommand()
	}
}

func (p *Parser) parseLCommand() string {
	command := p.currentCommand()
	return command[1 : len(command)-1]
}

func (p *Parser) currentCommand() string {
	if p.file != nil {
		return p.file[p.LineNumber]
	}
	return ""
}

func (p *Parser) commandType() types.Command {
	p.resetCommands()
	var commandType types.Command
	switch string(p.currentCommand()[0]) {
	case "@":
		commandType = types.ACommand
	case "(":
		commandType = types.LCommand
	default:
		commandType = types.CCommand
	}
	return commandType
}

func (p *Parser) parseACommand() string {
	return p.currentCommand()[1:]
}

func (p *Parser) Comp() (string, error) {
	if p.currentCommandType == types.CCommand {
		return p.comp, nil
	}
	return "", fmt.Errorf("comp command cannot be returned for %v, only C Commands", p.currentCommandType)
}

func (p *Parser) Dest() (string, error) {
	if p.currentCommandType == types.CCommand {
		return p.dest, nil
	}
	return "", fmt.Errorf("dest command cannot be returned for %v, only C Commands", p.currentCommandType)
}

func (p *Parser) Jump() (string, error) {
	if p.currentCommandType == types.CCommand {
		return p.jump, nil
	}
	return "", fmt.Errorf("jump command cannot be returned for %v, only c commands", p.currentCommandType)
}

func (p *Parser) Symbol() (string, error) {
	if p.currentCommandType != types.LCommand {
		return p.symbol, nil
	}
	return "", fmt.Errorf("symbol cannot be returned for %v, only c commands", p.currentCommandType)
}

func (p *Parser) parseCCommand() *CInstruction {
	return &CInstruction{}
}
