package main

import (
	"errors"
	"fmt"
)

// 一般情况下只有快速失败的过程，才会考虑panic
// 常用recover，从panic中恢复
func main() {
	//ErrorPkg()
	//errors.New("error")

	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello, panic: %v \n", data)
		}
		fmt.Println("恢复后从这里重新执行")
	}()

	panic("Boom！Boom！")
	fmt.Println("程序中止，不会执行到这里！")
}

func ErrorPkg() {
	err := &MyError{}
	// 使用%w 占位符，返回的是一个新错误
	// wrappedErr
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	// 再解出来
	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped")
		fmt.Println(err)
	}

	if errors.Is(wrappedErr, err) {
		// 虽然被包了一下，但是Is会逐层接触包装，判断是不是该err
		fmt.Println("wrapped is err")
	}

	copyErr := &MyError{}
	// 尝试将wrappedErr转换为MyError
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}

type MyError struct {
}

func (m *MyError) Error() string {
	return "is my error"
}
