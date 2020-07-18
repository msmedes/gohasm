package assembler

type Assembler struct {
	name string
}

func NewAssembler() *Assembler {
	return &Assembler{
		name: "hello",
	}
}
