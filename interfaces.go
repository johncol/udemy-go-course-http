package main

import "fmt"

// InterfaceA sample interface
type InterfaceA interface {
	methodA() string
}

// InterfaceB sample interface
type InterfaceB interface {
	methodB(int) string
}

// InterfaceC sample interface
type InterfaceC interface  {
	InterfaceA
	InterfaceB
}

func sampleCallerMethod(i InterfaceC) {
	fmt.Println(i.methodA())
	fmt.Println(i.methodB(0))
}

// MyStruct sample s truct
type MyStruct struct {
}

func (s MyStruct) methodA() string {
	return "MethodA string"
}

func (s MyStruct) methodB(number int) string {
	return "MethodB string"
}

func interfacesInheritanceSampleCode() {
	s := MyStruct{}
	sampleCallerMethod(s)
}