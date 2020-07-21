package parser

import (
	"bufio"
	"fmt"
	"gohasm/internal/pkg/types"
	"os"
	"strings"
)

type CInstruction struct {
	comp string
	dest string
	jump string
}

type Parser struct {
	filePath           string
	LineNumber         int
	file               []string
	CurrentCommandType types.Command
	symbol             string
	comp               string
	dest               string
	jump               string
	InstructionCounter int16
}

func New(filePath string) *Parser {
	return &Parser{
		filePath:           filePath,
		LineNumber:         -1,
		file:               loadFile(filePath),
		InstructionCounter: 0,
	}
}

func loadFile(filePath string) []string {
	var output []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = processLine(line)
		if line != "" {
			output = append(output, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return output
}

func processLine(line string) string {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "/") || line == "\n" {
		return ""
	}
	spaceIndex := strings.Index(line, " ")
	if spaceIndex != -1 {
		return line[:spaceIndex]
	}
	return line
}

func (p *Parser) Reset() {
	p.LineNumber = -1
}

func (p *Parser) resetCommands() {
	p.CurrentCommandType = types.NoCommand
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
	p.CurrentCommandType = p.CommandType()
	if p.CurrentCommandType == types.ACommand {
		p.symbol = p.parseACommand()
	}
	if p.CurrentCommandType == types.CCommand {
		instruction := p.parseCCommand()
		p.dest = instruction.dest
		p.comp = instruction.comp
		p.jump = instruction.jump
	}
	if p.CurrentCommandType == types.LCommand {
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

func (p *Parser) CommandType() types.Command {
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
	if p.CurrentCommandType == types.CCommand {
		return p.comp, nil
	}
	return "", fmt.Errorf("comp command cannot be returned for %v, only C Commands", p.CurrentCommandType)
}

func (p *Parser) Dest() (string, error) {
	if p.CurrentCommandType == types.CCommand {
		return p.dest, nil
	}
	return "", fmt.Errorf("dest command cannot be returned for %v, only C Commands", p.CurrentCommandType)
}

func (p *Parser) Jump() (string, error) {
	if p.CurrentCommandType == types.CCommand {
		return p.jump, nil
	}
	return "", fmt.Errorf("jump command cannot be returned for %v, only c commands", p.CurrentCommandType)
}

func (p *Parser) Symbol() (string, error) {
	if p.CurrentCommandType != types.CCommand {
		return p.symbol, nil
	}
	return "", fmt.Errorf("symbol cannot be returned for %v, only A or L commands", p.CurrentCommandType)
}

func (p *Parser) parseCCommand() *CInstruction {
	dest := "NONE"
	jump := "NONE"
	comp := ""
	equalIndex := strings.Index(p.currentCommand(), "=")
	if equalIndex != -1 {
		dest = p.currentCommand()[:equalIndex]
	}
	semiIndex := strings.Index(p.currentCommand(), ";")
	if semiIndex != -1 {
		jump = p.currentCommand()[semiIndex+1:]
		comp = p.currentCommand()[equalIndex+1 : semiIndex]
	} else {
		jump = "NONE"
		comp = p.currentCommand()[equalIndex+1:]
	}
	return &CInstruction{comp, dest, jump}
}
