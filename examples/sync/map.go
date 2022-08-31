package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("Cat", "Red")
	m.Store("Mouse", "Red")

	val, ok := m.Load("Cat")
	if ok {
		// 类型断言 t,ok := x.(T)	t := x.(T)
		// T可以是结构体或者指针
		// 类似Java instanceOf+强制类型转换合体
		// 如果是nil，那么永远是false
		//编译器不会检查
		fmt.Println(len(val.(string)))
	}
}
