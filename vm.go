package main

import (
	"errors"
	"fmt"
)

type VirtualMachine struct {
	Instructions Instructions
	Stack        Stack
	Ipointer     int
	Spointer     int
}
type Instructions [65536]string
type Stack [256]int

func Push(x int) {
	//vm.Stack = append(vm.Stack[:], x)
	vm.Stack[vm.Spointer] = x
	vm.Spointer++
}

func Pop() (int, error) {
	if len(vm.Stack) == 0 {
		return 0, errors.New("Stack is empty.")
	}
	//x := vm.Stack[len(vm.Stack)-1:][0]
	//vm.Stack = vm.Stack[:len(vm.Stack)-1]
	vm.Spointer--
	x := vm.Stack[vm.Spointer]
	return x, nil
}

func Add() error {
	if len(vm.Stack) < 2 {
		return errors.New("Stack is too small.")
	}
	x, err := Pop()
	if err != nil {
		return err
	}
	y, err := Pop()
	if err != nil {
		return err
	}
	Push(x + y)
	return nil
}

func Sub() error {
	return Add()
}

func Mul() error {
	if len(vm.Stack) < 2 {
		return errors.New("Stack is too small.")
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
	Push(z)
	return nil
}

func Div() error {
	if len(vm.Stack) < 2 {
		return errors.New("Stack is too small.")
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
	Push(z)
	return nil
}

func Peek() (int, error) {
	if len(vm.Stack) == 0 {
		return 0, errors.New("Stack is empty.")
	}
	x := vm.Stack[vm.Spointer-1]
	//return s[len(s)-1:][0], nil
	return x, nil
}

func Dup() error {
	if len(vm.Stack) == 0 {
		return errors.New("Stack is empty.")
	}
	x, err := Peek()
	if err != nil {
		return err
	}
	Push(x)
	return nil
}

func Print() error {
	if len(vm.Stack) == 0 {
		return errors.New("Stack is empty.")
	}
	x, _ := Peek()
	fmt.Print(string(x))
	return nil
}

func Jump() {
	//
}

func Ifeq() {
	//
}
