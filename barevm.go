// Package barevm implements a simple, bare virtual machine.
package barevm

import (
	"fmt"
)

// VM instructions
const (
	PSH = iota // Push a value onto the stack
	ADD        // Add two values pushed onto the stack
	MUL        // Multiply two values pushed onto the stack
	POP        // Pop a value of the stack
	HLT        // Halt program execution
)

// General purpose VM registers
const (
	A = iota
	B
	C
	D
	E
	F
	IP // Instruction pointer
	SP // Stack pointer
)

const (
	instructionCount = 5
	registerCount    = 8
	stackSize        = 256
)

// Program contains VM instruction.
type Program []int

// BareVM instance
type BareVM struct {
	program   []int
	registers [registerCount]int
	stack     [stackSize]int
	running   bool
}

// New creates a new instance of the VM
func New(program Program) *BareVM {
	p := make(Program, len(program))
	copy(p, program)
	return &BareVM{program: p}
}

// Run begins the execution of the program. Stops when it encounters a HLT
// instruction.
func (b *BareVM) Run() {
	b.running = true
	for b.running {
		i := b.fetch()
		b.eval(i)
		b.registers[IP]++
	}
}

// fetch fetches the next instruction from the program.
func (b *BareVM) fetch() int {
	return b.program[b.registers[IP]]
}

// eval evaluates an instruction.
func (b *BareVM) eval(i int) {
	switch i {
	case HLT:
		b.running = false
		break
	case PSH:
		b.registers[IP]++
		b.stack[b.registers[SP]] = b.program[b.registers[IP]]
		b.registers[SP]++
		break
	case POP:
		pop := b.stack[b.registers[SP]]
		b.registers[SP]--
		fmt.Println("pop:", pop)
		break
	case ADD:
		b.registers[SP]--
		x := b.stack[b.registers[SP]]

		b.registers[SP]--
		y := b.stack[b.registers[SP]]

		result := y + x

		b.registers[SP]++
		b.stack[b.registers[SP]] = result
		break
	case MUL:
		b.registers[SP]--
		x := b.stack[b.registers[SP]]

		b.registers[SP]--
		y := b.stack[b.registers[SP]]

		result := y * x

		b.registers[SP]++
		b.stack[b.registers[SP]] = result
		break
	}
}
