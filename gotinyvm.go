package main

import (
	"bufio"
	//"errors"
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
	err := runProgram()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nOK.")
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
	vm.InstructionCount = i
}

func runProgram() error {
	for {
		if vm.Ipointer >= vm.InstructionCount {
			return nil
			//if vm.Ipointer < 0 || vm.Ipointer >= vm.InstructionCount {
			//return errors.New("Invalid instruction address.")
		}
		if vm.Instructions[vm.Ipointer] == "" {
			//return errors.New("Invalid instruction at" + string(vm.Ipointer))
			//break
			//fmt.Println("skipping")
			vm.Ipointer++
			continue
		}
		instr := string(vm.Instructions[vm.Ipointer])
		instrTokens := strings.Split(instr, " ")
		op := instrTokens[0]
		var val string
		if len(instrTokens) > 1 {
			val = instrTokens[1]
		}

		switch op {
		case "push":
			v, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			err = Push(v)
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "pop":
			_, err := Pop()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "add":
			err := Add()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "sub":
			err := Sub()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "mul":
			err := Mul()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "div":
			err := Div()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "peek":
			_, err := Peek()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "dup":
			err := Dup()
			if err != nil {
				return err
			}
			vm.Ipointer++
		case "print":
			Print()
			vm.Ipointer++
		}
	}
	return nil
}
