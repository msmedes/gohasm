package symbol

// Table is a struct that contains a symbol table and counter for setting new memory addresses
type Table struct {
	table   map[string]int16
	counter int16
}

// New returns a new Table
func New() *Table {
	return &Table{
		table: map[string]int16{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"r0":     0,
			"R0":     0,
			"r1":     1,
			"R1":     1,
			"r2":     2,
			"R2":     2,
			"r3":     3,
			"R3":     3,
			"r4":     4,
			"R4":     4,
			"r5":     5,
			"R5":     5,
			"r6":     6,
			"R6":     6,
			"r7":     7,
			"R7":     7,
			"r8":     8,
			"R8":     8,
			"r9":     9,
			"R9":     9,
			"r10":    10,
			"R10":    10,
			"r11":    11,
			"R11":    11,
			"r12":    12,
			"R12":    12,
			"r13":    13,
			"R13":    13,
			"r14":    14,
			"R14":    14,
			"r15":    15,
			"R15":    15,
			"SCREEN": 16834,
			"KBD":    24576,
		},
		counter: 16,
	}
}

// AddEntry adds a new symbol when the address is known
func (st *Table) AddEntry(symbol string, addr int16) {
	st.table[symbol] = addr
}

// AddVariable adds a new symbol to the symbol table using the
// next available memory address, starting at 16.
func (st *Table) AddVariable(symbol string) (int16, bool) {
	st.table[symbol] = st.counter
	st.counter++
	addr, ok := st.table[symbol]
	return addr, ok
}

// GetAddr is a getter for addresses
func (st *Table) GetAddr(symbol string) (int16, bool) {
	addr, ok := st.table[symbol]
	return addr, ok
}

// Contains returns whether or not the symbol table contains
// the given symbol
func (st *Table) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}
