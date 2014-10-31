package main

import (
	"bufio"
	"fmt"
	"os"
)

type ProgramListing [65536]string

const (
	opPush  = 0x01
	opPop   = 0x02
	opJump  = 0x03
	opAdd   = 0x04
	opSub   = 0x05
	opMul   = 0x06
	opDiv   = 0x07
	opLoad  = 0x08
	opStore = 0x09
	opPeek  = 0x0a
	opDup   = 0x0b
	opPrint = 0x0c
)

var vm VirtualMachine
var pl ProgramListing
var plInstrCount int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotinyvm progname.bc")
		os.Exit(1)
	}
	vm = VirtualMachine{}
	vm.Stack = Stack{}
	loadProgram(os.Args[1])
	asm := NewAssembler()
	vm.Instructions = asm.Assemble(pl, plInstrCount)
	err := runProgram()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nOK.")
}

// loadProgram reads the program file and stores it as the ProgramListing.
// This function also sets the vm instruction count and ProgramListing instruction count.
func loadProgram(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to read program file,", path)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		//vm.Instructions[i] = scanner.Text()
		pl[i] = scanner.Text()
		i++
	}
	vm.InstructionCount = i
	plInstrCount = i
}

func runProgram() error {
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
			err := Push(v)
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opPop:
			_, err := Pop()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opJump:
			ival := int(val)
			vm.Ipointer = ival
		case opAdd:
			err := Add()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opSub:
			err := Sub()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opMul:
			err := Mul()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opDiv:
			err := Div()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opLoad:

		case opPeek:
			_, err := Peek()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opDup:
			err := Dup()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case opPrint:
			Print()
			vm.Ipointer++
		}
	}
	return nil
}
