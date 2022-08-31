package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
}

// 声明为全局变量
var once sync.Once

// PrintOnce 这个方法不管调用几次，只会输出一次
func PrintOnce() {
	once.Do(func() {
		fmt.Println("只会输出一次")
	})
}
