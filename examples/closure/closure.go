package main

import "fmt"

func main() {
	i := 1
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()
	fmt.Println(ReturnClosure("tom")())

	Delay()
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello," + name
	}
}

// Delay 延迟绑定
func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("this is: %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
