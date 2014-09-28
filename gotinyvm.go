package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var vm VirtualMachine

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vmtiny progname.bc")
		os.Exit(1)
	}
	vm = VirtualMachine{}
	vm.Stack = Stack{}
	loadProgram(os.Args[1])
	runProgram()
}

func loadProgram(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to read program file,", path)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		vm.Instructions[i] = scanner.Text()
		i++
	}
}

func runProgram() {
	for i := range vm.Instructions {
		if vm.Instructions[i] == "" {
			break
		}
		instr := strings.Split(string(vm.Instructions[i]), " ")
		op := instr[0]
		var val string
		if len(instr) > 1 {
			val = instr[1]
		}

		switch op {
		case "push":
			v, _ := strconv.Atoi(val)
			Push(v)
		case "print":
			Print()
		}
	}
}
