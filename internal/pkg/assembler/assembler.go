package assembler

import (
	"bufio"
	"fmt"
	"gohasm/internal/pkg/code"
	"gohasm/internal/pkg/parser"
	"gohasm/internal/pkg/symbol"
	"gohasm/internal/pkg/types"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Assembler struct {
	filePath  string
	parser    *parser.Parser
	code      *code.Code
	symbol    *symbol.Table
	Buffer    []string
	assembled bool
}

func New(filePath string) *Assembler {
	return &Assembler{
		filePath:  filePath,
		parser:    parser.New(filePath),
		code:      code.New(),
		symbol:    symbol.New(),
		Buffer:    make([]string, 0),
		assembled: false,
	}
}

func (a *Assembler) Assemble() {
	a.processLCommands()
	for a.parser.HasMoreCommands() {
		a.parser.Advance()
		commandType := a.parser.CurrentCommandType
		if commandType == types.ACommand {
			a.processACommand()
		}
		if commandType == types.CCommand {
			a.processCCommand()
		}
	}
	a.assembled = true
}

func (a *Assembler) processLCommands() {
	for a.parser.HasMoreCommands() {
		a.parser.Advance()
		if a.parser.CurrentCommandType == types.LCommand {
			sym, err := a.parser.Symbol()
			if err != nil {
				fmt.Println(err)
			}
			if !a.symbol.Contains(sym) {
				a.symbol.AddEntry(sym, a.parser.InstructionCounter)
			}
		} else {
			a.parser.InstructionCounter++
		}
	}
	a.parser.Reset()
}

func (a *Assembler) processACommand() {
	// Is there a more idiomatic way to do this because its gross
	sym, err := a.parser.Symbol()
	if err != nil {
		log.Fatal(err)
	}
	addr, err := strconv.Atoi(sym)
	if err == nil {
		binary := fmt.Sprintf("%016b", addr)
		a.Buffer = append(a.Buffer, binary)
	} else {
		addr, ok := a.symbol.GetAddr(sym)
		if !ok {
			addr, ok = a.symbol.AddVariable(sym)
		}
		binary := fmt.Sprintf("%016b", addr)
		a.Buffer = append(a.Buffer, binary)
	}

}

func (a *Assembler) processCCommand() {
	parserDest, err := a.parser.Dest()
	if err != nil {
		log.Fatal("parser does not have a destination for C Command")
	}
	parserComp, err := a.parser.Comp()
	if err != nil {
		log.Fatal("parser does not have a comparator for C Command")
	}
	parserJump, err := a.parser.Jump()
	if err != nil {
		log.Fatal("parser does not have a jump for C Command")
	}
	dest, ok := a.code.Dest(parserDest)
	if !ok {
		log.Fatalf("destination %v has no code", parserDest)
	}
	comp, ok := a.code.Comp(parserComp)
	if !ok {
		log.Fatalf("comparator %v has no code", parserComp)
	}
	jump, ok := a.code.Jump(parserJump)
	if !ok {
		log.Fatalf("jump %v has no code", parserJump)
	}
	a.Buffer = append(a.Buffer, fmt.Sprintf("111%s%s%s", comp, dest, jump))
}

func (a *Assembler) WriteToFile() {
	if a.assembled {
		dir, file := filepath.Split(a.filePath)
		extensionIndex := strings.Index(file, ".")
		fileName := file[:extensionIndex]
		hackFileName := fmt.Sprintf("%v/%v.hack", dir, fileName)
		writeFile, err := os.OpenFile(hackFileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(writeFile)

		for _, line := range a.Buffer {
			_, _ = datawriter.WriteString(line + "\n")
		}

		datawriter.Flush()
		writeFile.Close()
	}
}
