package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assembler struct {
	//
}

type ByteInstruction struct {
	ByteCode byte
	InstrVal byte
}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (a *Assembler) Assemble(assemblyInstructions ProgramListing, lineCount int) [65536]ByteInstruction {
	machineInstructions := [65536]ByteInstruction{}
	for i := 0; i < lineCount; i++ {
		instr := string(assemblyInstructions[i])
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

		switch instrName {
		case "push":
			byteCode := byte(0x01)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "pop":
			byteCode := byte(0x02)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "jump":
			byteCode := byte(0x03)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "add":
			byteCode := byte(0x04)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "sub":
			byteCode := byte(0x05)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "mul":
			byteCode := byte(0x06)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "div":
			byteCode := byte(0x07)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "load":
			byteCode := byte(0x08)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "store":
			byteCode := byte(0x09)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "peek":
			byteCode := byte(0x0a)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "dup":
			byteCode := byte(0x0b)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		case "print":
			byteCode := byte(0x0c)
			machineInstructions[i] = ByteInstruction{byteCode, instrVal}
		}
	}
	return machineInstructions
}
