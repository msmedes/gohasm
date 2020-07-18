package code

import "fmt"

// Code is a struct containing tables for comparison, destinations, and jump commands.
type Code struct {
	compTable map[string]string
	destTable map[string]string
	jumpTable map[string]string
}

// NewCode returns a Code struct
func NewCode() *Code {
	return &Code{
		compTable: map[string]string{
			"0":   "0101010",
			"1":   "0111111",
			"-1":  "0111010",
			"D":   "0001100",
			"A":   "0110000",
			"M":   "1110000",
			"!D":  "0001101",
			"!A":  "0110001",
			"!M":  "1110001",
			"-D":  "0001111",
			"-A":  "0110011",
			"-M":  "1110011",
			"D+1": "0011111",
			"A+1": "0110111",
			"M+1": "1110111",
			"D-1": "0001110",
			"A-1": "0110010",
			"M-1": "1110010",
			"D+A": "0000010",
			"D+M": "1000010",
			"D-A": "0010011",
			"D-M": "1010011",
			"A-D": "0000111",
			"M-D": "1000111",
			"D&A": "0000000",
			"D&M": "1000000",
			"D|A": "0010101",
			"D|M": "1010101",
		},
		destTable: map[string]string{
			"NONE": "000",
			"M":    "001",
			"D":    "010",
			"MD":   "011",
			"A":    "100",
			"AM":   "101",
			"AD":   "110",
			"AMD":  "111",
		},
		jumpTable: map[string]string{
			"NONE": "000",
			"JGT":  "001",
			"JEQ":  "010",
			"JGE":  "011",
			"JLT":  "100",
			"JNE":  "101",
			"JLE":  "110",
			"JMP":  "111",
		},
	}
}

// Comp is a getter for comparison operators
func (c *Code) Comp(mnem string) (string, error) {
	binary, ok := c.compTable[mnem]
	if !ok {
		return "", fmt.Errorf("Mnemonic %v is not in the comparison table", mnem)
	}
	return binary, nil
}

// Jump is a getter for jump operators
func (c *Code) Jump(mnem string) (string, error) {
	binary, ok := c.jumpTable[mnem]
	if !ok {
		return "", fmt.Errorf("Mnemonic %v is not in the jump table", mnem)
	}
	return binary, nil
}

// Dest is a getter for destination operators
func (c *Code) Dest(mnem string) (string, error) {
	binary, ok := c.destTable[mnem]
	if !ok {
		return "", fmt.Errorf("Mnemonic %v is not in the destination table", mnem)
	}
	return binary, nil
}
