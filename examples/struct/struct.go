package main

import "fmt"

type ToyDuck struct {
	Color string
	Price int
}

func (t ToyDuck) Swim() {
	fmt.Println("~")
}

func main() {
	duck1 := &ToyDuck{} // 推荐写法
	duck1.Swim()

	duck2 := ToyDuck{} // 推荐写法
	duck2.Swim()

	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你这样声明的时候，Go帮你分配好指针，ToyDuck = ToyDuck{}
	// 不用担心空指针，因为他就不是指针
	var duck4 ToyDuck
	duck4.Swim()

	// duck5就是一个指针，但是不知道指向哪里
	var duck5 *ToyDuck
	// 这边会直接panic
	duck5.Swim()

	// 赋值
	duck6 := ToyDuck{
		Color: "绿色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不推荐使用，耦合太高
	duck7 := ToyDuck{"蓝色", 10}
	duck7.Swim()

	// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "红色"
}
