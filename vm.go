package main

import (
	"errors"
	"fmt"
)

// VirtualMachine represents a virtual machine.
type VirtualMachine struct {
	Instructions     Instructions
	Stack            Stack
	Ipointer         int
	Spointer         int
	Registers        Registers
	InstructionCount int
}

// Instructions is a list of instructions received from the assembler.
type Instructions [65536]ByteInstruction
type Stack [256]int
type Registers [16]int

func NewVirtualMachine() *VirtualMachine {
	vm := VirtualMachine{}
	vm.Stack = Stack{}
	return &vm
}

// runProgram iterates the vm instructions received from the assembler
// and makes the appropriate calls.
func (vm *VirtualMachine) runProgram() error {
	for {
		if vm.Ipointer >= vm.InstructionCount {
			return nil
			//if vm.Ipointer < 0 || vm.Ipointer >= vm.InstructionCount {
			//return errors.New("Invalid instruction address.")
		}
		instr := vm.Instructions[vm.Ipointer]
		op := instr.ByteCode
		val := instr.InstrVal

		switch op {
		case opPush:
			v := int(val)
			err := vm.Push(v)
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opPop:
			_, err := vm.Pop()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opJump:
			ival := int(val)
			vm.Ipointer = ival
		case opAdd:
			err := vm.Add()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opSub:
			err := vm.Sub()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opMul:
			err := vm.Mul()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opDiv:
			err := vm.Div()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opLoad:

		case opPeek:
			_, err := vm.Peek()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opDup:
			err := vm.Dup()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opPrint:
			vm.Print()
			vm.Ipointer++
		}
	}
	return nil
}

func (vm *VirtualMachine) Push(x int) error {
	if vm.Spointer == 256 {
		return errors.New("Stack overflow.")
	}
	vm.Stack[vm.Spointer] = x
	vm.Spointer++
	return nil
}

func (vm *VirtualMachine) Pop() (int, error) {
	if vm.Spointer == 0 {
		return 0, errors.New("Stack underflow.")
	}
	vm.Spointer--
	x := vm.Stack[vm.Spointer]
	return x, nil
}

func (vm *VirtualMachine) Add() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := vm.Pop()
	if err != nil {
		return err
	}
	y, err := vm.Pop()
	if err != nil {
		return err
	}
	err = vm.Push(x + y)
	if err != nil {
		return err
	}
	return nil
}

func (vm *VirtualMachine) Sub() error {
	return vm.Add()
}

func (vm *VirtualMachine) Mul() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := vm.Pop()
	if err != nil {
		return err
	}
	y, err := vm.Pop()
	if err != nil {
		return err
	}
	z := 0
	for i := 0; i < y; i++ {
		z = z + x
	}
	err = vm.Push(z)
	if err != nil {
		return err
	}
	return nil
}

func (vm *VirtualMachine) Div() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := vm.Pop()
	if err != nil {
		return err
	}
	y, err := vm.Pop()
	if err != nil {
		return err
	}
	z := 0
	for x >= 0 {
		x = x - y
		z++
	}
	err = vm.Push(z)
	if err != nil {
		return err
	}
	return nil
}

func (vm *VirtualMachine) Peek() (int, error) {
	if vm.Spointer == 0 {
		return 0, errors.New("Stack underflow.")
	}
	x := vm.Stack[vm.Spointer-1]
	return x, nil
}

func (vm *VirtualMachine) Load(registerIdx int) error {
	if vm.Spointer == 256 {
		return errors.New("Stack overflow.")
	}
	err := vm.Push(vm.Registers[registerIdx])
	if err != nil {
		return err
	}
	return nil
}

func (vm *VirtualMachine) Store(registerIdx int) error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	v, err := vm.Pop()
	if err != nil {
		return err
	}
	vm.Registers[registerIdx] = v
	return nil
}

func (vm *VirtualMachine) Dup() error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	x, err := vm.Peek()
	if err != nil {
		return err
	}
	err = vm.Push(x)
	if err != nil {
		return err
	}
	return nil
}

func (vm *VirtualMachine) Print() error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	x, err := vm.Peek()
	if err != nil {
		return err
	}
	fmt.Print(string(x))
	return nil
}

func (vm *VirtualMachine) Ifeq() {
	//
}
