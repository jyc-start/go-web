package main

import "fmt"

func main() {
	// 创建一个预估容量为2的map
	m := make(map[string]string, 2)
	// 没有指定预估容量
	m1 := make(map[string]string)
	// 直接初始化
	m2 := map[string]string{
		"Tom": "111",
	}

	// 赋值
	m["hello"] = "world"
	m1["hello"] = "world"
	m2["hello"] = "world"

	// 取值
	val := m["hello"]
	println(val)

	val, ok := m["not_key"]
	if !ok {
		println("key not found")
	}
	// map遍历顺序不确定
	for k, v := range m {
		fmt.Printf("key:%s value:%s \n", k, v)
	}
}
