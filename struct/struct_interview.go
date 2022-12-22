package main

import "fmt"

type student struct {
	name string
	age  int
}

// 结构体初始化指针指向
func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "大王八", age: 9000},
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
	}
	for _, stu := range stus {
		fmt.Println(&stu.name)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(&v.name)
		fmt.Println(k, "=>", v.name)
	}
}
