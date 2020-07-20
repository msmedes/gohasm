package main

import (
	"gohasm/internal/pkg/assembler"
	"os"
)

func main() {
	fileName := os.Args[1]
	a := assembler.New(fileName)
	a.Assemble()
	a.WriteToFile()
}
