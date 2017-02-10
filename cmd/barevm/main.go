package main

import (
	"github.com/umahmood/barevm"
)

func main() {
	program := barevm.Program{
		barevm.PSH, 5,
		barevm.PSH, 6,
		barevm.ADD,
		barevm.POP,
		barevm.PSH, 42,
		barevm.PSH, 2,
		barevm.MUL,
		barevm.POP,
		barevm.HLT,
	}
	vm := barevm.New(program)
	vm.Run()
}
