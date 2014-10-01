package main

import (
	"errors"
	"fmt"
)

type VirtualMachine struct {
	Instructions     Instructions
	Stack            Stack
	Ipointer         int
	Spointer         int
	Registers        Registers
	InstructionCount int
}
type Instructions [65536]string
type Stack [256]int
type Registers [16]int

func Push(x int) error {
	if vm.Spointer == 256 {
		return errors.New("Stack overflow.")
	}
	vm.Stack[vm.Spointer] = x
	vm.Spointer++
	return nil
}

func Pop() (int, error) {
	if vm.Spointer == 0 {
		return 0, errors.New("Stack underflow.")
	}
	vm.Spointer--
	x := vm.Stack[vm.Spointer]
	return x, nil
}

func Add() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := Pop()
	if err != nil {
		return err
	}
	y, err := Pop()
	if err != nil {
		return err
	}
	err = Push(x + y)
	if err != nil {
		return err
	}
	return nil
}

func Sub() error {
	return Add()
}

func Mul() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := Pop()
	if err != nil {
		return err
	}
	y, err := Pop()
	if err != nil {
		return err
	}
	z := 0
	for i := 0; i < y; i++ {
		z = z + x
	}
	err = Push(z)
	if err != nil {
		return err
	}
	return nil
}

func Div() error {
	if vm.Spointer < 2 {
		return errors.New("Stack underflow.")
	}
	x, err := Pop()
	if err != nil {
		return err
	}
	y, err := Pop()
	if err != nil {
		return err
	}
	z := 0
	for x >= 0 {
		x = x - y
		z++
	}
	err = Push(z)
	if err != nil {
		return err
	}
	return nil
}

func Peek() (int, error) {
	if vm.Spointer == 0 {
		return 0, errors.New("Stack underflow.")
	}
	x := vm.Stack[vm.Spointer-1]
	return x, nil
}

func Load(registerIdx int) error {
	if vm.Spointer == 256 {
		return errors.New("Stack overflow.")
	}
	err := Push(vm.Registers[registerIdx])
	if err != nil {
		return err
	}
	return nil
}

func Store(registerIdx int) error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	v, err := Pop()
	if err != nil {
		return err
	}
	vm.Registers[registerIdx] = v
	return nil
}

func Dup() error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	x, err := Peek()
	if err != nil {
		return err
	}
	err = Push(x)
	if err != nil {
		return err
	}
	return nil
}

func Print() error {
	if vm.Spointer == 0 {
		return errors.New("Stack underflow.")
	}
	x, err := Peek()
	if err != nil {
		return err
	}
	fmt.Print(string(x))
	return nil
}

func Ifeq() {
	//
}
