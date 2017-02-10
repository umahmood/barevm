# Bare VM

Bare VM as the name suggests is a bare bones implementation of a virtual machine. 
It is very very basic, and was written just to show the reader the basic skeleton 
of a virtual machine.

# Installation

> $ go get github.com/umahmood/barevm

# Usage

```
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
```
Output:
```
pop 11
pop 84
```

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
