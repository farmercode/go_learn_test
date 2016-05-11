package main

import "fmt"

func main() {
	//常量定义
	const Name = "King"
	const Age = 35
	fmt.Printf("Name:%s\tAge:%d\n",Name,Age)

	//枚举类型
	const (
		Jan = iota + 1
		Feb = iota + 1
		Mar
	)
	fmt.Printf("%d\n%d\n%d\n",Jan, Feb, Mar)
}
