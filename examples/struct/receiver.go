package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// ChangeName 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// ChangeAge 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

// 遇事不决选指针！
func main() {
	user := &User{}
	user.ChangeName("明")
	user.ChangeAge(10)
	fmt.Println(user.Age, user.Name) // 10
}

// ?
type Handle func()

func (h Handle) Hello() {

}
