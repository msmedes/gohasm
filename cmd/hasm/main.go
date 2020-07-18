package main

import (
	"fmt"
	"gohasm/internal/pkg/assembler"
)

func main() {
	fmt.Println("hello world")
	a := assembler.NewAssembler()
	fmt.Printf("%+v", a)
}