package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Assembler represents the assembler.
type Assembler struct {
	//
}

// ByteInstruction is an opcode and value pair.
type ByteInstruction struct {
	ByteCode byte
	InstrVal byte
}

// NewAssembler returns a new Assembler.
func NewAssembler() *Assembler {
	return &Assembler{}
}

// Assemble iterates the lines of the program listing and converts them to bytecode
// instructions consisting of an opcode and a value argument for the opcode call.
func (a *Assembler) Assemble(statements ProgramListing, lineCount int) [65536]ByteInstruction {
	machineInstructions := [65536]ByteInstruction{}
	for i := 0; i < lineCount; i++ {
		instr := string(statements[i])
		instrTokens := strings.Split(instr, " ")
		instrName := instrTokens[0]
		var instrVal byte
		if len(instrTokens) > 1 {
			ival, err := strconv.Atoi(instrTokens[1])
			if err != nil {
				fmt.Println("Assembly error on line", i+1, err)
				os.Exit(1)
			}
			instrVal = byte(ival)
		}

		// TODO: decide if to write concisely, or leave if needed for later features.
		switch instrName {
		case "push":
			byteCode := byte(opPush)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "pop":
			byteCode := byte(opPop)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "jump":
			byteCode := byte(opJump)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "add":
			byteCode := byte(opAdd)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "sub":
			byteCode := byte(opSub)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "mul":
			byteCode := byte(opMul)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "div":
			byteCode := byte(opDiv)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "load":
			byteCode := byte(opLoad)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "store":
			byteCode := byte(opStore)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "peek":
			byteCode := byte(opPeek)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "dup":
			byteCode := byte(opDup)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "print":
			byteCode := byte(opPrint)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		}
	}
	return machineInstructions
}
