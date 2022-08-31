package main

import "fmt"

// go没有重写
type Parent struct {
}

func (p Parent) SayHello() {
	fmt.Println("i am " + p.Name())
}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

func (s Son) Name() string {
	return "son"
}

func main() {
	son := &Son{}
	son.SayHello() // i am Parent
}
